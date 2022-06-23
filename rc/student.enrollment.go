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

func getStudentEnrollment(ctx *gin.Context) {
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sid, _, err := getStudentRecruitmentCycleID(ctx, rid)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	var result []getStudentEnrollmentResponse

	err = fetchStudentQuestionsAnswers(ctx, rid, sid, &result)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, result)
}

func postEnrollmentAnswer(ctx *gin.Context) {
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var answer RecruitmentCycleQuestionsAnswer

	err = ctx.ShouldBindJSON(&answer)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	srcid, verified, err := getStudentRecruitmentCycleID(ctx, rid)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	if verified {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Already Verified"})
		return
	}

	answer.StudentRecruitmentCycleID = srcid

	err = createStudentAnswer(ctx, &answer)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": fmt.Sprintf("Answer %d created", answer.ID)})
}
