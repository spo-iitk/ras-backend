package student

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func deleteStudentHandler(ctx *gin.Context) {

	sid, err := strconv.ParseUint(ctx.Param("sid"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = deleteStudent(ctx, uint(sid))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("A student with id %d is deleted", sid)

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully deleted"})

}
