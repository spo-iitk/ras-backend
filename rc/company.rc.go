package rc

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/util"
)

type companyRecruitmentCycleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func getCompanyRecruitmentCycle(ctx *gin.Context) {

	var rcs []RecruitmentCycle
	companyID, err := extractCompanyID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	err = fetchRCsByCompanyID(ctx, companyID, &rcs)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	var rcsr []companyRecruitmentCycleResponse
	for _, rc := range rcs {
		rcsr = append(rcsr, companyRecruitmentCycleResponse{ID: rc.ID, Name: string(rc.Type) + " " + rc.AcademicYear})
	}
	ctx.JSON(200, rcsr)
}

type companyRCHRResponse struct {
	Name string `json:"name"`
	HR1  string `json:"hr1"`
	HR2  string `json:"hr2"`
	HR3  string `json:"hr3"`
}

func getCompanyRCHRHandler(ctx *gin.Context) {
	companyID, err := extractCompanyID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
	}

	var company CompanyRecruitmentCycle
	err = fetchCompanyByRCIDAndCID(ctx, util.ParseString(rid), util.ParseString(companyID), &company)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, companyRCHRResponse{
		Name: company.CompanyName,
		HR1:  company.HR1,
		HR2:  company.HR2,
		HR3:  company.HR3,
	})
}
