package student

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func updateStudentHandler(ctx *gin.Context) {
	var updateStudentRequest Student

	if err := ctx.ShouldBindJSON(&updateStudentRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := updateStudent(ctx, &updateStudentRequest, uint(id))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !updated {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	logrus.Infof("A student with id %d is updated", id)

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully updated"})
}