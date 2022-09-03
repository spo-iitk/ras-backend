package application

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/spo-iitk/ras-backend/rc"
	"github.com/spo-iitk/ras-backend/util"
	"google.golang.org/api/calendar/v3"
)

type getAllEventsByRCResponse struct {
	ProformaEvent
	CompanyName string `json:"company_name"`
	Role        string `json:"role"`
	Profile     string `json:"profile"`
}

func getAllEventsByRCHandler(ctx *gin.Context) {
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var events []getAllEventsByRCResponse

	err = fetchEventsByRC(ctx, rid, &events)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func getEventsByProformaHandler(ctx *gin.Context) {

	pid, err := util.ParseUint(ctx.Param("pid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var events []ProformaEvent
	err = fetchEventsByProforma(ctx, pid, &events)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func postEventHandler(ctx *gin.Context) {
	pid, err := util.ParseUint(ctx.Param("pid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var event ProformaEvent
	err = ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ProformaID = pid

	err = createEvent(ctx, &event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func putEventHandler(ctx *gin.Context) {
	var event ProformaEvent
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if event.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "event id is required"})
		return
	}

	err = updateEvent(ctx, &event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if event.StartTime == 0 && event.EndTime == 0 {
		ctx.JSON(http.StatusOK, event)
		return
	}

	var proforma Proforma

	err = fetchProforma(ctx, event.ProformaID, &proforma)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	time_zone := "Asia/Kolkata"
	loc, _ := time.LoadLocation(time_zone)

	rc.CreateNotice(ctx, rid, &rc.Notice{
		Title: fmt.Sprintf("%s of profile %s - %s has been scheduled", event.Name, proforma.Profile, proforma.CompanyName),
		Description: fmt.Sprintf(
			"%s of profile %s - %s has been scheduled from %s to %s",
			event.Name, proforma.Profile, proforma.CompanyName,
			time.UnixMilli(int64(event.StartTime)).In(loc).Format("2006-01-02 15:04"),
			time.UnixMilli(int64(event.EndTime)).In(loc).Format("2006-01-02 15:04")),
		Tags: fmt.Sprintf("scheduled,%s,%s,%s,%d", event.Name, proforma.Role, proforma.CompanyName, event.ID),
	})

	var cID string
	if proforma.RecruitmentCycleID == 1 {
		cID = viper.GetString("CALENDAR.CID1")
	}

	if proforma.RecruitmentCycleID == 2 {
		cID = viper.GetString("CALENDAR.CID2")
	}

	if cID == "" {
		ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Please as web head to generate a new calender in admin.events:144"})
		return
	}

	cevent := &calendar.Event{
		Summary:  fmt.Sprintf("%s of profile %s - %s", event.Name, proforma.Profile, proforma.CompanyName),
		Location: event.Venue,
		Description: fmt.Sprintf(
			"%s of profile %s - %s has been scheduled from %s to %s\nhttps://placement.iitk.ac.in/student/rc/%d/event/%d",
			event.Name, proforma.Profile, proforma.CompanyName,
			time.UnixMilli(int64(event.StartTime)).In(loc).Format("2006-01-02 15:04"),
			time.UnixMilli(int64(event.EndTime)).In(loc).Format("2006-01-02 15:04"),
			proforma.RecruitmentCycleID, event.ID),
		Start: &calendar.EventDateTime{
			DateTime: fmt.Sprintf("%s+05:30", time.UnixMilli(int64(event.StartTime)).In(loc).Format("2006-01-02T15:04")),
			TimeZone: time_zone,
		},
		End: &calendar.EventDateTime{
			DateTime: fmt.Sprintf("%s+05:30", time.UnixMilli(int64(event.StartTime)).In(loc).Format("2006-01-02T15:04")),
			TimeZone: time_zone,
		},
	}

	if event.CalID == "" {
		cevent, err = cal_srv.Events.Insert(cID, cevent).Do()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		event.CalID = cevent.Id

		err = updateEvent(ctx, &event)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	cevent, err = cal_srv.Events.Update(cID, event.CalID, cevent).Do()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func deleteEventHandler(ctx *gin.Context) {
	eid, err := util.ParseUint(ctx.Param("eid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var cevent ProformaEvent
	err = fetchEvent(ctx, eid, &cevent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rid := ctx.Param("rid")

	var cID string
	if rid == "1" {
		cID = viper.GetString("CALENDAR.CID1")
	}

	if rid == "2" {
		cID = viper.GetString("CALENDAR.CID2")
	}

	if cID == "" {
		ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Please as web head to generate a new calender in admin.events:218"})
		return
	}

	err = deleteEvent(ctx, eid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = cal_srv.Events.Delete(cID, cevent.CalID).Do()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully deleted event"})
}
