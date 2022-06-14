package application

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/util"
)

func getProformaForCompanyHandler(ctx *gin.Context) {
	cid, err := extractCompanyRCID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	var jps []Proforma

	err = fetchProformaByCompanyRC(ctx, cid, &jps)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, jps)
}

func postProformaByCompanyHandler(ctx *gin.Context) {
	cid, err := extractCompanyRCID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	companyid, err := extractCompanyID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	var jp Proforma

	err = ctx.ShouldBindJSON(&jp)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	jp.CompanyRecruitmentCycleID = cid
	jp.CompanyID = companyid
	jp.IsApproved = sql.NullBool{}

	err = createProforma(ctx, &jp)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v created a proforma with id %d", user, jp.ID)
	ctx.JSON(200, gin.H{"pid": jp.ID})
}

func putProformaByCompanyHandler(ctx *gin.Context) {
	cid, err := extractCompanyRCID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	var jp Proforma
	err = ctx.ShouldBindJSON(&jp)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	jp.CompanyRecruitmentCycleID = cid

	err = updateProformaForCompany(ctx, &jp)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "edited proforma"})
}

func deleteProformaByCompanyHandler(ctx *gin.Context) {
	pid, err := util.ParseUint(ctx.Param("pid"))
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	cid, err := extractCompanyRCID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	ok, err := deleteProformaByCompany(ctx, pid, cid)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		ctx.AbortWithStatusJSON(402, gin.H{"error": "proforma not found or unauthorized"})
		return
	}

	ctx.JSON(200, gin.H{"status": "deleted proforma"})
}
