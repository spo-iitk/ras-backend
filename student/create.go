package student

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func createStudentHandler(ctx *gin.Context) {
	var createStudentRequest Student

	if err := ctx.ShouldBindJSON(&createStudentRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := createStudent(ctx, &createStudentRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("A student %s is created with id %d", createStudentRequest.Name, createStudentRequest.ID)

	ctx.JSON(http.StatusOK, gin.H{"status": "Success"})

}
