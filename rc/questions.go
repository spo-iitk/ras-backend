package rc

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/middleware"
)

func getStudentQuestions(ctx *gin.Context) {
	rid := ctx.Param("rid")
	var questions []RecruitmentCycleQuestion

	err := fetchStudentQuestions(ctx, rid, &questions)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, questions)
}

func postStudentQuestion(ctx *gin.Context) {
	var question RecruitmentCycleQuestion

	err := ctx.BindJSON(&question)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = createStudentQuestion(ctx, &question)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	qid := question.ID
	user := middleware.GetUserID(ctx)

	logrus.Infof("%v created a student question with id %d", user, qid)

	ctx.JSON(200, gin.H{"data": qid})
}

func putStudentQuestion(ctx *gin.Context) {
	var question RecruitmentCycleQuestion

	err := ctx.BindJSON(&question)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = updateStudentQuestion(ctx, &question)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v updated a student question with id %d", user, question.ID)

	ctx.JSON(200, gin.H{"status": "updated student question"})
}

func deleteStudentQuestionHandler(ctx *gin.Context) {
	qid := ctx.Param("qid")

	err := deleteStudentQuestion(ctx, qid)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v deleted a student question with id %d", user, qid)

	ctx.JSON(200, gin.H{"status": "deleted student question"})
}
