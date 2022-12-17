package rc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/util"
)

type companyRecruitmentCycleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func getCompanyRCHandler(ctx *gin.Context) {
	var rcs []RecruitmentCycle

	companyID, err := extractCompanyID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = fetchRCsByCompanyID(ctx, companyID, &rcs)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var rcsr []companyRecruitmentCycleResponse
	for _, rc := range rcs {
		rcsr = append(rcsr, companyRecruitmentCycleResponse{ID: rc.ID, Name: string(rc.Type) + " " + rc.AcademicYear})
	}

	ctx.JSON(http.StatusOK, rcsr)
}

type getAllRCHandlerForCompanyResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Enrolled bool   `json:"enrolled"`
}

func getAllRCHandlerForCompany(ctx *gin.Context) {
	companyID, err := extractCompanyID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var enrolledRcs []RecruitmentCycle
	err = fetchRCsByCompanyID(ctx, companyID, &enrolledRcs)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var allRcs []RecruitmentCycle
	err = fetchAllActiveRCs(ctx, &allRcs)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var enrolledRCMap = make(map[uint]RecruitmentCycle)
	for _, rc := range enrolledRcs {
		enrolledRCMap[rc.ID] = rc
	}

	var allRCMap = make(map[uint]RecruitmentCycle)
	for _, rc := range allRcs {
		allRCMap[rc.ID] = rc
	}

	var companyAllRCResponse []getAllRCHandlerForCompanyResponse
	for _, rc := range allRCMap {
		isEnrolled := false
		if _, ok := enrolledRCMap[rc.ID]; ok {
			isEnrolled = true
		}

		companyAllRCResponse = append(companyAllRCResponse, getAllRCHandlerForCompanyResponse{
			ID:       rc.ID,
			Name:     string(rc.Type) + " " + rc.AcademicYear,
			Enrolled: isEnrolled,
		})
	}

	ctx.JSON(http.StatusOK, companyAllRCResponse)
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var company CompanyRecruitmentCycle
	err = fetchCompanyByRCIDAndCID(ctx, rid, companyID, &company)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, companyRCHRResponse{
		Name: company.CompanyName,
		HR1:  company.HR1,
		HR2:  company.HR2,
		HR3:  company.HR3,
	})
}

type EnrollCompanyRequest struct {
	CompanyName string `json:"company_name" binding:"required"`
	HR1         string `json:"hr1" binding:"required"`
	HR2         string `json:"hr2"`
	HR3         string `json:"hr3"`
}

func enrollCompanyHandler(ctx *gin.Context) {
	var enrollCompany EnrollCompanyRequest

	err := ctx.ShouldBindJSON(&enrollCompany)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	companyID, err := extractCompanyID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var company = CompanyRecruitmentCycle{
		CompanyID:          companyID,
		CompanyName:        enrollCompany.CompanyName,
		RecruitmentCycleID: uint(rid),
		HR1:                enrollCompany.HR1,
		HR2:                enrollCompany.HR2,
		HR3:                enrollCompany.HR3,
	}

	err = upsertCompany(ctx, &company)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, company)
}
