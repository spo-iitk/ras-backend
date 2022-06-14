package student

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/middleware"
)

func getStudentHandler(ctx *gin.Context) {
	var student Student
	email := middleware.GetUserID(ctx)

	err := getStudentByEmail(ctx, &student, email)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, student)

}
