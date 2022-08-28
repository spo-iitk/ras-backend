package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/util"
)

func getQuestionsByProformaHandler(ctx *gin.Context) {
	pid, err := util.ParseUint(ctx.Param("pid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rcid, err := util.ParseUint(ctx.Param("rcid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var questions []ApplicationQuestion
	err = fetchProformaQuestion(ctx, pid, &questions)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i := 0; i < len(questions); i++ {
		var answer ApplicationQuestionAnswer
		err = fetchProformaQuestionAnswer(ctx, questions[i].QuestionID, rcid, &answer)

	}
	ctx.JSON(http.StatusOK, questions)
}

func postQuestionHandler(ctx *gin.Context) {
	var question ApplicationQuestion

	err := ctx.ShouldBindJSON(&question)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if question.Question == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "question is required"})
		return
	}

	pid, err := util.ParseUint(ctx.Param("pid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	question.ProformaID = pid

	if question.Type == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "type is required"})
		return
	}

	err = createProformaQuestion(ctx, &question)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v created a proforma question with id %d", user, question.ID)
	ctx.JSON(http.StatusOK, gin.H{"qid": question.ID})
}

func putQuestionHandler(ctx *gin.Context) {
	var question ApplicationQuestion

	err := ctx.ShouldBindJSON(&question)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if question.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	err = updateProformaQuestion(ctx, &question)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v updated a proforma question with id %d", user, question.ID)

	ctx.JSON(http.StatusOK, gin.H{"status": "updated question successfully"})
}

func deleteQuestionHandler(ctx *gin.Context) {
	qid, err := util.ParseUint(ctx.Param("qid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = deleteProformaQuestion(ctx, qid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v deleted a proforma question with id %d", user, qid)

	ctx.JSON(http.StatusOK, gin.H{"status": "deleted question successfully"})
}

func fetchQuestionAnswers(ctx *gin.Context, qid uint, rcid uint) {
	rcid, err := util.ParseUint(ctx.Param("rcid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var answer ApplicationQuestionAnswer
	err = fetchProformaQuestionAnswer(ctx, qid, rcid, &answer)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
}
