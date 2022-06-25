package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/util"
)

func getProformasForStudentHandler(ctx *gin.Context) {
	_, err := extractStudentRCID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Fail to extract student rcid"})
		return
	}

	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var jps []Proforma

	err = fetchProformasForStudent(ctx, rid, &jps)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, jps)
}

func getProformaForStudentHandler(ctx *gin.Context) {
	_, err := extractStudentRCID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Fail to extract student rcid"})
		return
	}

	pid, err := util.ParseUint(ctx.Param("pid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var jp Proforma

	err = fetchProformaForStudent(ctx, pid, &jp)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, jp)
}
