package rc

import (
	"net/http"
	"strconv"

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

	err := ctx.ShouldBindJSON(&question)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	rid, err := strconv.ParseUint(ctx.Param("rid"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	question.RecruitmentCycleID = uint(rid)
	err = createStudentQuestion(ctx, &question)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v created a student question with id %d", user, question.ID)

	ctx.JSON(200, question)
}

func putStudentQuestion(ctx *gin.Context) {
	var question RecruitmentCycleQuestion

	err := ctx.ShouldBindJSON(&question)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	if question.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Enter question ID"})
		return
	}

	ok, err := updateStudentQuestion(ctx, &question)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No such question exists"})
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
