package rc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/util"
)

type AllResumeResponse struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	Sid           uint   `json:"sid"`
	RsId          uint   `json:"rsid"`
	Resume        string `json:"resume"`
	Verified      bool   `json:"verified"`
	ActionTakenBy string `json:"action_taken_by"`
}

func getAllResumes(ctx *gin.Context) {
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resumes []AllResumeResponse
	err = fetchAllResumes(ctx, rid, &resumes)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resumes)
}

func getResume(ctx *gin.Context) {
	rsid, err := util.ParseUint(ctx.Param("rsid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resume, err := fetchResume(ctx, rsid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resume)
}

type putResumeVerifyRequest struct {
	Verified bool `json:"verified"`
}

func putResumeVerify(ctx *gin.Context) {
	rsid, err := util.ParseUint(ctx.Param("rsid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var req putResumeVerifyRequest

	err = ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ok, err := updateResumeVerify(ctx, rsid, req.Verified)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "resume not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}
