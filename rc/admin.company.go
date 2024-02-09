package rc

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/util"
)

func getAllCompaniesHandler(ctx *gin.Context) {
	rid := ctx.Param("rid")
	var companies []CompanyRecruitmentCycle

	err := fetchAllCompanies(ctx, rid, &companies)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, companies)
}

type addNewCompanyRequest struct {
	CompanyID   uint   `gorm:"index" json:"company_id" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	HR1         string `json:"hr1" binding:"required"`
	HR2         string `json:"hr2"`
	HR3         string `json:"hr3"`
	Comments    string `json:"comments"`
}

func postNewCompanyHandler(ctx *gin.Context) {
	var addNewCompany addNewCompanyRequest

	err := ctx.ShouldBindJSON(&addNewCompany)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if addNewCompany.CompanyID == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "company_id is required"})
		return
	}

	rid, err := strconv.ParseUint(ctx.Param("rid"), 10, 32)
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

	err = upsertCompany(ctx, &company)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, company)
}

func putCompanyHandler(ctx *gin.Context) {
	var editCompanyRequest CompanyRecruitmentCycle

	err := ctx.ShouldBindJSON(&editCompanyRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if editCompanyRequest.CompanyID != 0 || editCompanyRequest.RecruitmentCycleID != 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "company_id or rid is not allowed"})
		return
	}

	if editCompanyRequest.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	err = editCompany(ctx, &editCompanyRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, editCompanyRequest)
}

func getCompanyHandler(ctx *gin.Context) {
	cid, err := util.ParseUint(ctx.Param("cid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var company CompanyRecruitmentCycle

	err = FetchCompany(ctx, cid, &company)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, company)
}

func deleteCompanyHandler(ctx *gin.Context) {
	cid, err := util.ParseUint(ctx.Param("cid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = deleteRCCompany(ctx, cid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v deleted %v from RC", user, cid)

	ctx.JSON(http.StatusOK, gin.H{"status": "Company Deleted from this RC"})
}

func getCompanyHistoryHandler(ctx *gin.Context) {
    var companyHistory []CompanyHistory
    cid, err := strconv.ParseUint(ctx.Param("cid"), 10, 32)
    if err != nil {
        ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err = FetchCompanyHistory(ctx, uint(cid), &companyHistory)
    if err != nil {
        ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var historyResponse []CompanyHistoryResponse
    for _, history := range companyHistory {
        historyResponse = append(historyResponse, CompanyHistoryResponse{
            ID:                 history.ID,
            RecruitmentCycleID:   history.RecruitmentCycleID,
            Comments:           history.Comments,
        })
    }

    ctx.JSON(http.StatusOK, historyResponse)
}