package rc

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/middleware"
)

func getStudentQuestionsHandler(ctx *gin.Context) {
	rid := ctx.Param("rid")
	var questions []RecruitmentCycleQuestion

	err := fetchStudentQuestions(ctx, rid, &questions)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, questions)
}

func postStudentQuestionHandler(ctx *gin.Context) {
	var question RecruitmentCycleQuestion

	err := ctx.ShouldBindJSON(&question)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rid, err := strconv.ParseUint(ctx.Param("rid"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	question.RecruitmentCycleID = uint(rid)
	err = createStudentQuestion(ctx, &question)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v created a student question with id %v", user, question.ID)

	ctx.JSON(http.StatusOK, question)
}

func putStudentQuestionHandler(ctx *gin.Context) {
	var question RecruitmentCycleQuestion

	err := ctx.ShouldBindJSON(&question)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if question.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Enter question ID"})
		return
	}

	ok, err := updateStudentQuestion(ctx, &question)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No such question exists"})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v updated a student question with id %d", user, question.ID)

	ctx.JSON(http.StatusOK, gin.H{"status": "updated student question"})
}

func deleteStudentQuestionHandler(ctx *gin.Context) {
	qid := ctx.Param("qid")

	err := deleteStudentQuestion(ctx, qid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v deleted a student question with id %v", user, qid)

	ctx.JSON(http.StatusOK, gin.H{"status": "deleted student question"})
}
