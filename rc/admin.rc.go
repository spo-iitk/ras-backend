package rc

import (
	"github.com/gin-gonic/gin"
)

func getAllRC(ctx *gin.Context) {
	var rc []RecruitmentCycle
	err := fetchAllRCs(ctx, &rc)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, rc)
}

type RC struct {
	IsActive            bool                 `json:"is_active" gorm:"default:true"`
	AcademicYear        string               `json:"academic_year" binding:"required"`
	Type                RecruitmentCycleType `json:"type" binding:"required"`
	StartDate           int64                `json:"start_date" binding:"required"`
	Phase               string               `json:"phase" binding:"required"`
	ApplicationCountCap uint                 `json:"application_count_cap" binding:"required"`
}

func postRC(ctx *gin.Context) {
	var recruitmentCycle RC
	err := ctx.ShouldBindJSON(&recruitmentCycle)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
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
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, gin.H{"id": rc.ID})
}

//! TODO: Add more response data
type getRCResponse struct {
	RecruitmentCycle
}

func getRC(ctx *gin.Context) {
	id := ctx.Param("rid")
	var rc RecruitmentCycle
	err := fetchRC(ctx, id, &rc)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, getRCResponse{rc})
}

func GetMaxCountfromRC(ctx *gin.Context) (uint, error) {
	id := ctx.Param("rid")
	var rc RecruitmentCycle
	err := fetchRC(ctx, id, &rc)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return 0, err
	}

	MaxCount := rc.ApplicationCountCap

	return MaxCount, nil
}

type editRCRequest struct {
	ID                  uint `json:"id" binding:"required"`
	Inactive            bool `json:"inactive" binding:"required"`
	ApplicationCountCap uint `json:"application_count_cap"`
}

func editRCHandler(ctx *gin.Context) {
	var req editRCRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	ok, err := updateRC(ctx, req.ID, req.Inactive, req.ApplicationCountCap)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Could not find data"})
		return
	}

	ctx.JSON(200, gin.H{"status": "Updated Succesfully"})
}
