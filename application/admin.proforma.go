package application

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/rc"
	"github.com/spo-iitk/ras-backend/util"
)

func getProformaHandler(ctx *gin.Context) {
	pid, err := util.ParseUint(ctx.Param("pid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var jp Proforma

	err = fetchProforma(ctx, pid, &jp)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, jp)
}

func getProformaByCompanyHandler(ctx *gin.Context) {
	cid, err := util.ParseUint(ctx.Param("cid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var jps []Proforma

	err = fetchProformasByCompanyForAdmin(ctx, cid, &jps)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, jps)
}

func getAllProformasHandler(ctx *gin.Context) {
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var jps []Proforma

	err = fetchProformaByRCForAdmin(ctx, rid, &jps)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, jps)
}

func putProformaHandler(ctx *gin.Context) {
	var jp Proforma

	err := ctx.ShouldBindJSON(&jp)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if jp.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	var oldJp Proforma
	err = fetchProforma(ctx, jp.ID, &oldJp)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jp.ActionTakenBy = middleware.GetUserID(ctx)

	publishNotice := oldJp.Deadline == 0 && jp.Deadline != 0

	err = updateProforma(ctx, &jp)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v edited a proforma with id %d", user, jp.ID)

	if publishNotice {
		logrus.Infof("%v published a proforma with id %d", user, jp.ID)

		err = rc.CreateNotice(ctx, oldJp.RecruitmentCycleID, &rc.Notice{
			Title: fmt.Sprintf("[%s] | New Job Opening for %s", jp.CompanyName, jp.Profile),
			Description: fmt.Sprintf(
				"A new opening has been created for the profile of %s in the company %s",
				jp.Profile, jp.CompanyName),
			Tags:     fmt.Sprintf("opening,%s,%s", jp.Role, jp.CompanyName),
			Deadline: jp.Deadline,
		})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "Proforma with id " + util.ParseString(jp.ID) + " has been published"})
	} else {

		ctx.JSON(http.StatusOK, gin.H{"status": "Updated proforma with id " + util.ParseString(jp.ID)})
	}
}

type hideProformaRequest struct {
	ID          uint `binding:"required"`
	HideDetails bool `json:"hide_details"`
}

func hideProformaHandler(ctx *gin.Context) {
	var req hideProformaRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = updateHideProforma(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v edited a proforma with id %d", user, req.ID)

	ctx.JSON(http.StatusOK, gin.H{"status": "Updated proforma with id " + util.ParseString(req.ID)})
}

func postProformaHandler(ctx *gin.Context) {
	var jp Proforma

	err := ctx.ShouldBindJSON(&jp)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = createProforma(ctx, &jp)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v created a proforma with id %d", user, jp.ID)

	ctx.JSON(http.StatusOK, gin.H{"pid": jp.ID})
}

func deleteProformaHandler(ctx *gin.Context) {
	pid, err := util.ParseUint(ctx.Param("pid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = deleteProforma(ctx, pid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "deleted proforma"})
}
