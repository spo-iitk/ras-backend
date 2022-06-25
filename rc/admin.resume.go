package rc

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/util"
)

type AllResumeResponse struct {
	Name          string       `json:"name"`
	Email         string       `json:"email"`
	Sid           uint         `json:"sid"`
	Rsid          uint         `json:"rsid"`
	Resume        string       `json:"resume"`
	Verified      sql.NullBool `json:"verified"`
	ActionTakenBy string       `json:"action_taken_by"`
}

func getAllResumesHandler(ctx *gin.Context) {
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var resumes []AllResumeResponse
	err = fetchAllResumes(ctx, rid, &resumes)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resumes)
}

func getResumeHandler(ctx *gin.Context) {
	rsid, err := util.ParseUint(ctx.Param("rsid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resume, err := FetchResume(ctx, rsid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resume)
}

type putResumeVerifyRequest struct {
	Verified bool `json:"verified"`
}

func putResumeVerifyHandler(ctx *gin.Context) {
	rsid, err := util.ParseUint(ctx.Param("rsid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req putResumeVerifyRequest

	err = ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	ok, err := updateResumeVerify(ctx, rsid, req.Verified, user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "resume not found"})
		return
	}

	logrus.Infof("%v verified resume with id %d, changed state to %v", user, rsid, req.Verified)

	ctx.JSON(http.StatusOK, gin.H{"status": true})
}
