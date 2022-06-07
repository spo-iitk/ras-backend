package rc

import (
	"github.com/gin-gonic/gin"
)

func getStudentEnrollment(ctx *gin.Context) {
	rid := ctx.Param("rid")
	sid := ctx.Param("sid")

	var questions []RecruitmentCycleQuestion
	var answers []RecruitmentCycleQuestionsAnswer

	err := fetchStudentQuestions(ctx, rid, &questions)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = fetchStudentAnswers(ctx, sid, &answers)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"questions": questions, "answers": answers})
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
