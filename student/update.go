package student

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/middleware"
)

func updateStudentByIDHandler(ctx *gin.Context) {
	var updateStudentRequest Student

	if err := ctx.ShouldBindJSON(&updateStudentRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updateStudentRequest.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Enter student ID"})
		return
	}

	updated, err := updateStudentByID(ctx, &updateStudentRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !updated {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	logrus.Infof("A student with id %d is updated", updateStudentRequest.ID)

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully updated"})
}

func updateStudentHandler(ctx *gin.Context) {
	var updateStudentRequest Student

	if err := ctx.ShouldBindJSON(&updateStudentRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := middleware.GetUserID(ctx)

	updated, err := updateStudentByEmail(ctx, &updateStudentRequest, email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !updated {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	logrus.Infof("A student with email %s is updated", email)

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully updated"})
}
