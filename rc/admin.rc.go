package rc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllRCHandler(ctx *gin.Context) {
	var rc []RecruitmentCycle
	err := fetchAllRCs(ctx, &rc)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if(ctx.GetInt("roleID") > 101){
		var activeRC []RecruitmentCycle
		for _, element := range rc {
			if element.IsActive {
				activeRC = append(activeRC, element)
			}
		}
		rc = activeRC
	}
	
	ctx.JSON(http.StatusOK, rc)
}

type RC struct {
	IsActive            bool                 `json:"is_active" gorm:"default:true"`
	AcademicYear        string               `json:"academic_year" binding:"required"`
	Type                RecruitmentCycleType `json:"type" binding:"required"`
	StartDate           int64                `json:"start_date" binding:"required"`
	Phase               string               `json:"phase" binding:"required"`
	ApplicationCountCap uint                 `json:"application_count_cap" binding:"required"`
}

func postRCHandler(ctx *gin.Context) {
	var recruitmentCycle RC
	err := ctx.ShouldBindJSON(&recruitmentCycle)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var rc = RecruitmentCycle{
		IsActive:            recruitmentCycle.IsActive,
		AcademicYear:        recruitmentCycle.AcademicYear,
		Type:                recruitmentCycle.Type,
		StartDate:           recruitmentCycle.StartDate,
		Phase:               recruitmentCycle.Phase,
		ApplicationCountCap: recruitmentCycle.ApplicationCountCap,
	}

	err = createRC(ctx, &rc)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, gin.H{"id": rc.ID})
}

//! TODO: Add more response data
type getRCResponse struct {
	RecruitmentCycle
}

func getRCHandler(ctx *gin.Context) {
	id := ctx.Param("rid")
	var rc RecruitmentCycle
	err := fetchRC(ctx, id, &rc)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, getRCResponse{rc})
}

func GetMaxCountfromRC(ctx *gin.Context) (uint, error) {
	id := ctx.Param("rid")
	var rc RecruitmentCycle

	err := fetchRC(ctx, id, &rc)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0, err
	}

	return rc.ApplicationCountCap, nil
}

type editRCRequest struct {
	ID                  uint `json:"id" binding:"required"`
	Inactive            bool `json:"inactive"`
	ApplicationCountCap uint `json:"application_count_cap"`
}

func editRCHandler(ctx *gin.Context) {
	var req editRCRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok, err := updateRC(ctx, req.ID, req.Inactive, req.ApplicationCountCap)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Could not find data"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Updated Succesfully"})
}
