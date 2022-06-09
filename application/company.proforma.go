package application

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/middleware"
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

	err = ctx.BindJSON(&jp)
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
	err = ctx.BindJSON(&jp)
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

	ctx.JSON(200, gin.H{"data": "edited proforma"})
}
