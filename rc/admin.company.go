package rc

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/util"
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
	HR1         string `json:"hr1" binding:"required"`
	HR2         string `json:"hr2"`
	HR3         string `json:"hr3"`
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

	ctx.JSON(200, company)
}

func getCompany(ctx *gin.Context) {
	cid, err := util.ParseUint(ctx.Param("cid"))
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	var company CompanyRecruitmentCycle

	err = FetchCompany(ctx, cid, &company)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, company)
}
