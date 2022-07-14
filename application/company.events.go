package application

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/util"
)

func postEventByCompanyHandler(ctx *gin.Context) {
	var event ProformaEvent
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cid := getCompanyRCID(ctx)
	if cid != 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "could not get company rcid"})
		return
	}

	var jp Proforma
	err = fetchProforma(ctx, event.ProformaID, &jp)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if jp.CompanyRecruitmentCycleID != cid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "company not authorized"})
		return
	}

	err = createEvent(ctx, &event)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "event created with id " + fmt.Sprint(event.ID)})
}

func putEventByCompanyHandler(ctx *gin.Context) {
	var event ProformaEvent
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if event.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	var curr_event ProformaEvent
	err = fetchEvent(ctx, event.ID, &curr_event)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cid := getCompanyRCID(ctx)
	if cid != 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "could not get company rcid"})
		return
	}

	var jp Proforma
	err = fetchProforma(ctx, curr_event.ProformaID, &jp)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if jp.CompanyRecruitmentCycleID != cid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "company not authorized"})
		return
	}

	err = updateEvent(ctx, &event)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "event created with id " + fmt.Sprint(event.ID)})
}

func deleteEventByCompanyHandler(ctx *gin.Context) {
	eid, err := util.ParseUint(ctx.Param("eid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cid := getCompanyRCID(ctx)
	if cid != 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "could not get company rcid"})
		return
	}

	var event ProformaEvent
	err = fetchEvent(ctx, eid, &event)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var jp Proforma
	err = fetchProforma(ctx, event.ProformaID, &jp)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if jp.CompanyRecruitmentCycleID != cid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "company not authorized"})
		return
	}

	err = deleteEvent(ctx, eid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "deleted event with id " + fmt.Sprint(eid)})
}

func getEventsByProformaForCompanyHandler(ctx *gin.Context) {
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
