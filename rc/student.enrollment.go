package rc

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/util"
)

type getStudentEnrollmentResponse struct {
	RecruitmentCycleQuestion
	Answer string `json:"answer"`
}

func getStudentEnrollmentHandler(ctx *gin.Context) {
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sid := getStudentRCID(ctx)
	if sid == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "SRCID not found"})
		return
	}

	var result []getStudentEnrollmentResponse

	err = fetchStudentQuestionsAnswers(ctx, rid, sid, &result)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func postEnrollmentAnswerHandler(ctx *gin.Context) {
	var answer RecruitmentCycleQuestionsAnswer

	err := ctx.ShouldBindJSON(&answer)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	srcid, verified, err := extractStudentRCID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if verified {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Already Verified"})
		return
	}

	answer.StudentRecruitmentCycleID = srcid

	err = createStudentAnswer(ctx, &answer)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("Answer %d created", answer.ID)})
}
