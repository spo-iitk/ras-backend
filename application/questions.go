package application

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/util"
)

func getQuestionsByPID(ctx *gin.Context) {
	pid_string := ctx.Param("pid")
	pid, err := util.ParseUint(pid_string)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	var questions []JobApplicationQuestion
	err = fetchPerformaQuestion(ctx, pid, &questions)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, questions)
}

func getQuestionsByQID(ctx *gin.Context) {
	qid_string := ctx.Param("qid")
	qid, err := util.ParseUint(qid_string)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	var question JobApplicationQuestion
	err = fetchPerformaQuestionByID(ctx, qid, &question)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, question)
}

func postQuestion(ctx *gin.Context) {
	var question JobApplicationQuestion
	err := ctx.BindJSON(&question)

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = createPerformaQuestion(ctx, &question)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v created a performa question with id %d", user, question.ID)

	ctx.JSON(200, gin.H{"qid": question.ID})
}

func putQuestion(ctx *gin.Context) {
	var question JobApplicationQuestion
	err := ctx.BindJSON(&question)

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = updatePerformaQuestion(ctx, &question)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v updated a performa question with id %d", user, question.ID)

	ctx.JSON(200, gin.H{"data": "updated question successfully"})
}
