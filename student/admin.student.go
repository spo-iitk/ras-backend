package student

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getStudentByIDHandler(ctx *gin.Context) {
	var student Student

	id, err := strconv.ParseUint(ctx.Param("sid"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = getStudentByID(ctx, &student, uint(id))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, student)
}

func getAllStudentsHandler(ctx *gin.Context) {
	var students []Student

	err := getAllStudents(ctx, &students)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, students)
}

func getLimitedStudentsHandler(ctx *gin.Context) {
	var students []Student

	pageSize := ctx.DefaultQuery("pageSize", "100")
	lastFetchedId := ctx.Query("lastFetchedId")
	batch := ctx.Query("batch")

	pageSizeInt, err := strconv.ParseInt(pageSize, 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lastFetchedIdInt, err := strconv.ParseInt(lastFetchedId, 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	batchInt, err := strconv.ParseInt(batch, 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = getLimitedStudents(ctx, &students, int(lastFetchedIdInt), int(pageSizeInt), int(batchInt))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, students)
}
