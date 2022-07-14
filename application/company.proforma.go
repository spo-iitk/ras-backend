package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/rc"
	"github.com/spo-iitk/ras-backend/util"
)

func getProformaForCompanyHandler(ctx *gin.Context) {
	cid := getCompanyRCID(ctx)
	if cid == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "could not get company rcid"})
		return
	}

	var jps []Proforma

	err := fetchProformasByCompanyForCompany(ctx, cid, &jps)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, jps)
}

func postProformaByCompanyHandler(ctx *gin.Context) {
	cid := getCompanyRCID(ctx)
	if cid == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "could not get company rcid"})
		return
	}

	var jp Proforma
	err := ctx.ShouldBindJSON(&jp)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var companyRC rc.CompanyRecruitmentCycle
	err = rc.FetchCompany(ctx, cid, &companyRC)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jp.CompanyRecruitmentCycleID = cid
	jp.CompanyID = companyRC.CompanyID
	jp.RecruitmentCycleID = companyRC.RecruitmentCycleID
	jp.CompanyName = companyRC.CompanyName

	err = createProforma(ctx, &jp)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v created a proforma with id %d", user, jp.ID)
	ctx.JSON(http.StatusOK, gin.H{"pid": jp.ID})
}

func putProformaByCompanyHandler(ctx *gin.Context) {
	cid := getCompanyRCID(ctx)
	if cid == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "could not get company rcid"})
		return
	}

	var jp Proforma
	err := ctx.ShouldBindJSON(&jp)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jp.CompanyRecruitmentCycleID = cid

	ok, err := updateProformaForCompany(ctx, &jp)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "proforma not found or unauthorized"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "edited proforma"})
}

func deleteProformaByCompanyHandler(ctx *gin.Context) {
	pid, err := util.ParseUint(ctx.Param("pid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cid := getCompanyRCID(ctx)
	if cid == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "could not get company rcid"})
		return
	}

	ok, err := deleteProformaByCompany(ctx, pid, cid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		ctx.AbortWithStatusJSON(402, gin.H{"error": "proforma not found or unauthorized"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "deleted proforma"})
}

func getProformaHandlerForCompany(ctx *gin.Context) {
	pid, err := util.ParseUint(ctx.Param("pid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cid := getCompanyRCID(ctx)
	if cid == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "could not get company rcid"})
		return
	}

	var jp Proforma
	err = fetchProformaForCompany(ctx, pid, cid, &jp)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, jp)
}
