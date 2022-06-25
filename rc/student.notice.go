package rc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllNoticesForStudentHandler(ctx *gin.Context) {
	_, _, err := extractStudentRCID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rid := ctx.Param("rid")
	var notices []Notice

	err = fetchAllNotices(ctx, rid, &notices)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, notices)
}
