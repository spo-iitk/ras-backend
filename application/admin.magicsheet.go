package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/util"
)

// GET /admin/magic-sheets/:rid
func getAllMagicSheetsHandler(ctx *gin.Context) {
	_, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid rc_id"})
		return
	}

	var sheets []MagicSheet
	err = FetchMagicSheetData(ctx, &sheets)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch magic sheets"})
		return
	}

	ctx.JSON(http.StatusOK, sheets)
}

// POST /admin/magic-sheets
func createMagicSheetHandler(ctx *gin.Context) {
	var sheet MagicSheet
	if err := ctx.ShouldBindJSON(&sheet); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := CreateMagicSheetData(ctx, &sheet)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, sheet)
}

// PUT /admin/magic-sheets/:id
func updateMagicSheetHandler(ctx *gin.Context) {
	var sheet MagicSheet
	if err := ctx.ShouldBindJSON(&sheet); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := UpdateMagicSheetFull(ctx, &sheet)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update magic sheet"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Updated"})
}

// DELETE /admin/magic-sheets/:id
func deleteMagicSheetHandler(ctx *gin.Context) {
	var sheet MagicSheet
	if err := ctx.ShouldBindJSON(&sheet); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	err := DeleteMagicSheetData(ctx, &sheet)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Could not delete magic sheet"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Deleted"})
}
