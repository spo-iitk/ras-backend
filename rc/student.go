package rc

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/middleware"
)

func getAllStudents(ctx *gin.Context) {
	rid := ctx.Param("rid")
	var students []StudentRecruitmentCycle

	err := fetchAllStudents(ctx, rid, &students)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, students)
}

func getStudent(ctx *gin.Context) {
	sid := ctx.Param("sid")
	var student StudentRecruitmentCycle

	err := fetchStudent(ctx, sid, &student)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, student)
}

func putStudent(ctx *gin.Context) {
	var student StudentRecruitmentCycle

	err := ctx.BindJSON(&student)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = updateStudent(ctx, &student)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v updated a student with id %d", user, student.ID)

	ctx.JSON(200, gin.H{"status": "updated student"})
}

func postStudents(ctx *gin.Context) {
	var students []StudentRecruitmentCycle

	err := ctx.BindJSON(&students)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = createStudents(ctx, &students)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)
	num := len(students)

	logrus.Infof("%v addedd %v new students", user, num)

	ctx.JSON(200, gin.H{"status": "created students", "data": students})
}
