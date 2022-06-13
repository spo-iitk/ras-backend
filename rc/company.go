package rc

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

type addNewCompanyRequest struct {
	CompanyID   uint   `gorm:"index" json:"company_id" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Comments    string `json:"comments"`
}

func postNewCompany(ctx *gin.Context) {
	var addNewCompany addNewCompanyRequest

	err := ctx.ShouldBindJSON(&addNewCompany)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	rid, err := strconv.ParseUint(ctx.Param("rid"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	var company = CompanyRecruitmentCycle{
		CompanyID:          addNewCompany.CompanyID,
		CompanyName:        addNewCompany.CompanyName,
		RecruitmentCycleID: uint(rid),
		Comments:           addNewCompany.Comments,
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
