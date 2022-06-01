package student

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllStudentsHandler(ctx *gin.Context) {
	var students []Student

	err := getAllStudents(ctx, &students)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": students})
}
