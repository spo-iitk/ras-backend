package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/util"
)

func getCoordinatorsByEventHandler(ctx *gin.Context) {
	eid_string := ctx.Param("eid")
	eid, err := util.ParseUint(eid_string)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	coordinators := []EventCoordinator{}
	err = fetchCoordinatorsByEvent(ctx, eid, &coordinators)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, coordinators)
}

type postCoordinatorByEventRequest struct {
	EventID       uint   `json:"event_id" binding:"required"`
	CoordinatorID string `json:"coordinator_id" binding:"required"`
	Name          string `json:"mame" binding:"required"`
}

func postCoordinatorByEventHandler(ctx *gin.Context) {
	var req postCoordinatorByEventRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var coordinator = EventCoordinator{
		ProformaEventID: req.EventID,
		CordinatorID:    req.CoordinatorID,
		CordinatorName:  req.Name,
	}

	err = createEventCoordinator(ctx, &coordinator)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Added successfully"})
}
