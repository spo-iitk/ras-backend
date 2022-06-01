package student

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getStudentHandler(ctx *gin.Context) {
	var student Student

	id, err := strconv.ParseUint(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = getStudent(ctx, &student, uint(id))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": student})
}
