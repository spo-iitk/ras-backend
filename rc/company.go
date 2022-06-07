package rc

import "github.com/gin-gonic/gin"

func getAllCompanies(ctx *gin.Context) {
	rid := ctx.Param("rid")
	var companies []CompanyRecruitmentCycle

	err := fetchAllCompanies(ctx, rid, &companies)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, companies)
}

func postNewCompany(ctx *gin.Context) {
	var company CompanyRecruitmentCycle

	err := ctx.BindJSON(&company)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = createCompany(ctx, &company)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	cid := company.ID
	ctx.JSON(200, gin.H{"data": cid})
}

func getCompany(ctx *gin.Context) {
	rid := ctx.Param("rid")
	cid := ctx.Param("cid")
	var company CompanyRecruitmentCycle

	err := fetchCompany(ctx, rid, cid, &company)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, company)
}
