package rc

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getRCCount(ctx *gin.Context) {
	var studentCount int
	var companyCount int

	rid, err := strconv.ParseUint(ctx.Param("rid"), 10, 64)
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

	ctx.JSON(200, gin.H{"registered_student": studentCount, "registered_company": companyCount})
}
