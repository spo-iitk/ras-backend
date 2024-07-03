package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/util"
)

func getPvfForVerificationHandler(ctx *gin.Context) {
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := util.ParseUint("1") // id to be upadated
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var jps PVF
	err = fetchPvfForVerification(ctx, id, rid, &jps)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, jps)
}

func verifyPvfHandler(ctx *gin.Context) {
	// var verifyPvfRequest PVF
	// if err := ctx.ShouldBindJSON(&verifyPvfRequest); err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// pid, err := util.ParseUint(ctx.Param("pid"))

	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// verifyPvfRequest.ID = pid
	// updated, err := verifyPvf(ctx, &verifyPvfRequest)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// if !updated {
	// 	ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "PVF not found"})
	// 	return
	// }

	// if verifyPvfRequest.IsVerified {
	// 	logrus.Infof("A PVF with id %d is verified", verifyPvfRequest.ID)
	// 	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully verified"})
	// } else {
	// 	logrus.Infof("A PVF with id %d is unverified", verifyPvfRequest.ID)
	// 	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully unverified"})
	// }

}

func putPVFHandler(ctx *gin.Context) {
	var jp PVF

	err := ctx.ShouldBindJSON(&jp)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if jp.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	var oldJp PVF
	err = fetchPVF(ctx, jp.ID, &oldJp)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// jp.ActionTakenBy = middleware.GetUserID(ctx)

	// publishNotice := oldJp.Deadline == 0 && jp.Deadline != 0

	err = updatePVF(ctx, &jp)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "Updated PVF with id " + util.ParseString(jp.ID)})
	// user := middleware.GetUserID(ctx)

	// logrus.Infof("%v edited a proforma with id %d", user, jp.ID)

	// if publishNotice {
	// 	logrus.Infof("%v published a proforma with id %d", user, jp.ID)

	// 	err = rc.CreateNotice(ctx, oldJp.RecruitmentCycleID, &rc.Notice{
	// 		Title: fmt.Sprintf("[%s] | New Job Opening for %s", jp.CompanyName, jp.Profile),
	// 		Description: fmt.Sprintf(
	// 			"A new opening has been created for the profile of %s in the company %s",
	// 			jp.Profile, jp.CompanyName),
	// 		Tags: fmt.Sprintf("opening,%s,%s", jp.Role, jp.CompanyName),
	// 	})
	// 	if err != nil {
	// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	ctx.JSON(http.StatusOK, gin.H{"status": "Proforma with id " + util.ParseString(jp.ID) + " has been published"})
	// } else {
	// 	ctx.JSON(http.StatusOK, gin.H{"status": "Updated proforma with id " + util.ParseString(jp.ID)})
	// }
}
