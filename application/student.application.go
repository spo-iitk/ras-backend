package application

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/rc"
	"github.com/spo-iitk/ras-backend/util"
)

func postApplicationHandler(ctx *gin.Context) {
	pid, err := util.ParseUint(ctx.Param("pid"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	eid, err := fetchApplicationEventID(ctx, pid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sid, err := extractStudentRCID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	proformaEligibility, err := getEligibility(ctx, pid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	eligible, err := rc.GetStudentEligible(ctx, sid, proformaEligibility)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !eligible {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not eligible to apply"})
		return
	}

	applicationCount, err := getCurrentApplicationCount(ctx, sid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	applicationMaxCount, err := rc.GetMaxCountfromRC(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if applicationCount >= int(applicationMaxCount) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Application count maxed out"})
		return
	}

	var application = EventStudent{
		ProformaEventID:           eid,
		StudentRecruitmentCycleID: sid,
		Present:                   true,
	}

	err = createEventStudent(ctx, &application)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("Application for %d submitted against Performa %d with application ID %s", sid, pid, application.ID)
	ctx.JSON(http.StatusOK, gin.H{"success": "application submitted with id: " + fmt.Sprint(application.ID)})
}

func deleteApplicationHandler(ctx *gin.Context) {
	pid, err := util.ParseUint(ctx.Param("pid"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sid, err := extractStudentRCID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = deleteApplication(ctx, pid, sid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("Application for %d deleted against Performa %d", sid, pid)
	ctx.JSON(http.StatusOK, gin.H{"success": "application deleted"})
}

func getEventHandler(ctx *gin.Context) {
	eid_string := ctx.Param("eid")
	eid, err := util.ParseUint(eid_string)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var event ProformaEvent
	err = fetchEvent(ctx, eid, &event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, event)
}
