package rc

import (
	"github.com/gin-gonic/gin"
)

func getStudentEnrollment(ctx *gin.Context) {
	rid := ctx.Param("rid")
	sid := ctx.Param("sid")
	var questions []RecruitmentCycleQuestion

	err := fetchStudentQuestions(ctx, rid, sid, &questions)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, questions)
}

func postEnrollmentAnswer(ctx *gin.Context) {
	var answer RecruitmentCycleQuestionsAnswer

	err := ctx.BindJSON(&answer)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = createStudentAnswer(ctx, &answer)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	aid := answer.ID
	ctx.JSON(200, gin.H{"data": aid})
}
