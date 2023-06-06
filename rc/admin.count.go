package rc

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getRCCountHandler(ctx *gin.Context) {
	var studentCount int
	var companyCount int

	rid, err := strconv.ParseUint(ctx.Param("rid"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	studentCount, err = getRegisteredStudentCount(ctx, uint(rid))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	companyCount, err = getRegisteredCompanyCount(ctx, uint(rid))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"registered_student": studentCount, "registered_company": companyCount})
}
