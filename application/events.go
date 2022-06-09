package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/rc"
	"github.com/spo-iitk/ras-backend/util"
)

func getAllEventsByRCHandler(ctx *gin.Context) {
	rid_string := ctx.Param("rid")
	rid, err := util.ParseUint(rid_string)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	events := []JobProformaEvent{}
	err = fetchEventsByRC(ctx, rid, &events)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func getEventsByPIDHandler(ctx *gin.Context) {
	pid_string := ctx.Param("pid")
	pid, err := util.ParseUint(pid_string)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	events := []JobProformaEvent{}
	err = fetchEventsByPID(ctx, pid, &events)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func postEventHandler(ctx *gin.Context) {
	pid_string := ctx.Param("pid")
	pid, err := util.ParseUint(pid_string)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var event JobProformaEvent
	err = ctx.BindJSON(&event)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.JobProformaID = pid

	err = createEvent(ctx, &event)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func putEventHandler(ctx *gin.Context) {
	var event JobProformaEvent
	err := ctx.BindJSON(&event)
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func deleteEventHandler(ctx *gin.Context) {
	eid_string := ctx.Param("eid")
	eid, err := util.ParseUint(eid_string)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = deleteEvent(ctx, eid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully deleted event"})
}

func getEventsByStudentHandler(ctx *gin.Context) {
	rid_string := ctx.Param("rid")
	rid, err := util.ParseUint(rid_string)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user_email := middleware.GetUserID(ctx)
	if user_email == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	sid, err := rc.FetchStudentRCIDs(ctx, rid, []string{user_email})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	events := []JobProformaEvent{}
	err = fetchEventsByStudent(ctx, sid[0], &events)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, events)
}
