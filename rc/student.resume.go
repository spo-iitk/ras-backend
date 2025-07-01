package rc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/util"
)

type ResumeRequest struct {
	Resume     string     `json:"resume"`
	ResumeType ResumeType `json:"resume_type"`
	ResumeTag  string     `json:"resume_tag"`
}

func postStudentResumeHandler(ctx *gin.Context) {
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var request ResumeRequest
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sid := getStudentRCID(ctx)
	if sid == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "SRCID not found"})
		return
	}

	err = addStudentResume(ctx, request.Resume, sid, rid, request.ResumeType, request.ResumeTag) // Include resumeType in the function call
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Resume Added Successfully"})
}

func getStudentResumeHandler(ctx *gin.Context) {
	sid := getStudentRCID(ctx)
	if sid == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "SRCID not found"})
		return
	}

	var resumes []StudentRecruitmentCycleResume
	err := fetchStudentResume(ctx, sid, &resumes)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resumes)
}
