package application

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

	var application = EventStudent{
		ProformaEventID:           eid,
		StudentRecruitmentCycleID: sid,
		Present:                   true,
	}
	var applications = []EventStudent{application}
	err = createEventStudents(ctx, &applications)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "application submitted with id: " + fmt.Sprint(applications[0].ID)})
}

func deleteApplicationHandler(ctx *gin.Context) {
	sid, err := extractStudentRCID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pid, err := util.ParseUint(ctx.Param("pid"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = deleteApplication(ctx, pid, sid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "application deleted"})
}

func getEventsByIDHandler(ctx *gin.Context) {
	eid_string := ctx.Param("eid")
	eid, err := util.ParseUint(eid_string)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var event ProformaEvent
	err = fetchEventByID(ctx, eid, &event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, event)
}
