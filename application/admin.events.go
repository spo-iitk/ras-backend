package application

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/rc"
	"github.com/spo-iitk/ras-backend/util"
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

	if (event.StartTime == 0 && event.EndTime == 0) || event.Name == string(ApplicationSubmitted) || event.Name == string(Recruited) || event.Name == string(PIOPPOACCEPTED) {
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

	if event.CalID == "" {
		var notice rc.Notice

		notice.Title = fmt.Sprintf("%s of profile %s - %s has been scheduled", event.Name, proforma.Profile, proforma.CompanyName)
		notice.Tags = fmt.Sprintf("scheduled,%s,%s,%s,%d", event.Name, proforma.Role, proforma.CompanyName, event.ID)

		notice.Description = fmt.Sprintf(
			"%s of profile %s - %s has been scheduled from %s to %s",
			event.Name, proforma.Profile, proforma.CompanyName,
			time.UnixMilli(int64(event.StartTime)).In(loc).Format("2006-01-02 15:04"),
			time.UnixMilli(int64(event.EndTime)).In(loc).Format("2006-01-02 15:04"))

		rc.CreateNotice(ctx, rid, &notice)
	}

	cID := getCalenderID(proforma.RecruitmentCycleID)

	if cID == "" {
		ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Please as web head to generate a new calender in admin.events:218"})
		return
	}

	ctx.JSON(http.StatusOK, event)

	go insertCalenderEvent(&event, &proforma, loc, time_zone, cID)
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

	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cID := getCalenderID(rid)

	if cID == "" {
		ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Please as web head to generate a new calender in admin.events:218"})
		return
	}

	err = deleteEvent(ctx, eid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully deleted event"})

	go deleteCalenderEvent(cID, &cevent)
}
