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

	if addNewCompany.CompanyID == 0 {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "company_id is required"})
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
		HR1:                addNewCompany.HR1,
		HR2:                addNewCompany.HR2,
		HR3:                addNewCompany.HR3,
		Comments:           addNewCompany.Comments,
	}

	err = createCompany(ctx, &company)

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, company)
}

func putNewCompany(ctx *gin.Context) {
	var editCompanyRequest CompanyRecruitmentCycle

	err := ctx.ShouldBindJSON(&editCompanyRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	if editCompanyRequest.CompanyID != 0 || editCompanyRequest.RecruitmentCycleID != 0 {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "company_id or rid is not allowed"})
		return
	}

	if editCompanyRequest.ID == 0 {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "id is required"})
		return
	}

	err = editCompany(ctx, &editCompanyRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, editCompanyRequest)
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
