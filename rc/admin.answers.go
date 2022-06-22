package rc

import "github.com/gin-gonic/gin"

func getStudentAnswers(ctx *gin.Context) {
	sid := ctx.Param("sid")
	var answers []RecruitmentCycleQuestionsAnswer

	err := fetchStudentAnswers(ctx, sid, &answers)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, answers)
}

// func putStudentAnswer(ctx *gin.Context) {
// 	var answer RecruitmentCycleQuestionsAnswer

// 	err := ctx.ShouldBindJSON(&answer)
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	err = updateStudentAnswer(ctx, &answer)
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	user := middleware.GetUserID(ctx)

// 	logrus.Infof("%v updated a student answer with id %d", user, answer.ID)

// 	ctx.JSON(200, gin.H{"status": "updated student answer"})
// }

// func deleteStudentAnswerHandler(ctx *gin.Context) {
// 	sid := ctx.Param("sid")
// 	qid := ctx.Param("qid")

// 	err := deleteStudentAnswer(ctx, qid, sid)
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	user := middleware.GetUserID(ctx)

// 	logrus.Infof("%v deleted a student answer with id %d", user, sid)

// 	ctx.JSON(200, gin.H{"status": "deleted student answer"})
// }
