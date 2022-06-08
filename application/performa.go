package application

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/util"
)

func getPerformaByCompanyID(ctx *gin.Context) {
	cid_string := ctx.Param("cid")
	cid, err := util.ParseUint(cid_string)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	var jps []JobProforma

	err = fetchPerformaByCompanyRC(ctx, cid, &jps)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, jps)
}

func getPerformaByRID(ctx *gin.Context) {
	rid_string := ctx.Param("rid")
	rid, err := util.ParseUint(rid_string)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	var jps []JobProforma

	err = fetchPerformaByRC(ctx, rid, &jps)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, jps)
}

func getPerformaByPID(ctx *gin.Context) {
	pid_string := ctx.Param("pid")
	pid, err := util.ParseUint(pid_string)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	var jp JobProforma

	err = fetchJobPerforma(ctx, pid, &jp)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, jp)
}

func postPerformaByCompanyID(ctx *gin.Context) {
	cid_string := ctx.Param("cid")
	cid, err := util.ParseUint(cid_string)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	var jp JobProforma

	err = ctx.BindJSON(&jp)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	if jp.CompanyRecruitmentCycleID != cid {
		ctx.AbortWithStatusJSON(500, gin.H{"error": "Company ID mismatch"})
		return
	}

	err = createJobPerforma(ctx, &jp)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v created a performa with id %d", user, jp.ID)

	ctx.JSON(200, gin.H{"pid": jp.ID})
}

func putPerforma(ctx *gin.Context) {
	var jp JobProforma

	err := ctx.BindJSON(&jp)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = updateJobPerforma(ctx, &jp)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v edited a performa with id %d", user, jp.ID)

	ctx.JSON(200, jp)
}

func putPerformaByCompanyID(ctx *gin.Context) {
	cid_string := ctx.Param("cid")
	cid, err := util.ParseUint(cid_string)

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	var jp JobProforma
	if jp.CompanyRecruitmentCycleID != cid {
		ctx.AbortWithStatusJSON(500, gin.H{"error": "Company ID mismatch"})
		return
	}

	err = ctx.BindJSON(&jp)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = updateJobPerforma(ctx, &jp)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": "edited performa"})
}

func deletePerformaByCompanyID(ctx *gin.Context) {
	pid_string := ctx.Param("pid")
	pid, err := util.ParseUint(pid_string)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	cid_string := ctx.Param("cid")
	cid, err := util.ParseUint(cid_string)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	var jp JobProforma
	err = fetchJobPerforma(ctx, pid, &jp)

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	if jp.CompanyRecruitmentCycleID != cid {
		ctx.AbortWithStatusJSON(500, gin.H{"error": "Company ID mismatch"})
		return
	}

	err = deleteJobPerforma(ctx, pid)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": "deleted performa"})
}
