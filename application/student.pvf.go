package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/util"
)

func postPvfForStudentHandler(ctx *gin.Context) {
	sid := getStudentRCID(ctx)
	if sid == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "SRCID not found"})
		return
	}
	var pvf PVF
	err := ctx.ShouldBindJSON(&pvf)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pvf.StudentRecruitmentCycleID = sid
	err = createPVF(ctx, &pvf)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := middleware.GetUserID(ctx)

	logrus.Infof("%v created \a pvf with id %d", user, pvf.ID)
	ctx.JSON(http.StatusOK, gin.H{"pid": pvf.ID})

}

func getAllPvfForStudentHandler(ctx *gin.Context) {
	sid := getStudentRCID(ctx)
	if sid == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "SRCID not found"})
		return
	}
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var jps []PVF
	err = fetchAllPvfForStudent(ctx, sid, rid, &jps)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, jps)

}
func getPvfForStudentHandler(ctx *gin.Context) {
	sid := getStudentRCID(ctx)
	if sid == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "SRCID not found"})
		return
	}
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pid, err := util.ParseUint(ctx.Param("pid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var jps PVF
	err = fetchPvfForStudent(ctx, sid, rid, pid, &jps)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, jps)
}
