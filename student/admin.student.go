package student

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/util"
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
	print(ctx.Request.URL.Query())
	pageSize := ctx.DefaultQuery("pageSize", "100")
	lastFetchedId := ctx.Query("lastFetchedId")
	batch := ctx.Query("batch")

	pageSizeInt, err := util.ParseUint(pageSize)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lastFetchedIdInt, err := util.ParseUint(lastFetchedId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	batchInt, err := util.ParseUint(batch)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = getLimitedStudents(ctx, &students, lastFetchedIdInt, pageSizeInt, batchInt)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, students)
}
