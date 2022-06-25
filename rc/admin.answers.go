package rc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/util"
)

func getStudentAnswersHandler(ctx *gin.Context) {
	sid, err := util.ParseUint(ctx.Param("sid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var answers []getStudentEnrollmentResponse

	err = fetchStudentQuestionsAnswers(ctx, rid, sid, &answers)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, answers)
}

// func putStudentAnswer(ctx *gin.Context) {
// 	var answer RecruitmentCycleQuestionsAnswer

// 	err := ctx.ShouldBindJSON(&answer)
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	err = updateStudentAnswer(ctx, &answer)
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	user := middleware.GetUserID(ctx)

// 	logrus.Infof("%v updated a student answer with id %d", user, answer.ID)

// 	ctx.JSON(http.StatusOK, gin.H{"status": "updated student answer"})
// }

// func deleteStudentAnswerHandler(ctx *gin.Context) {
// 	sid := ctx.Param("sid")
// 	qid := ctx.Param("qid")

// 	err := deleteStudentAnswer(ctx, qid, sid)
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	user := middleware.GetUserID(ctx)

// 	logrus.Infof("%v deleted a student answer with id %d", user, sid)

// 	ctx.JSON(http.StatusOK, gin.H{"status": "deleted student answer"})
// }
