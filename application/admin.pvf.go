package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/util"
)

func getAllPvfForAdminHandler(ctx *gin.Context) {
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var jps []PVF
	err = fetchAllPvfForAdmin(ctx, rid, &jps)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, jps)

}

func getPvfForAdminHandler(ctx *gin.Context) {
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
	err = fetchPvfForAdmin(ctx, rid, pid, &jps)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, jps)
}

func sendVerificationLinkForPvfHandler(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pid, err := util.ParseUint(ctx.Param("pid"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var pvf PVF

		rid, err := util.ParseUint(ctx.Param("rid"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = fetchPvfForAdmin(ctx, rid, pid, &pvf)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := middleware.GeneratePVFToken("akshat23@iitk.ac.in", pid, rid)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		logrus.Infof("A Token %s with  id %d", token, pid) // to be removed

		// hardcode email
		mail_channel <- mail.GenerateMail("iitkakshat@gmail.com",
			"Verification requested on RAS",
			"Dear "+pvf.MentorName+"PVF ID :"+util.ParseString(pid)+"Token :  "+token+",\n\nWe got your request for registration on Recruitment Automation System, IIT Kanpur. We will get back to you soon. For any queries, please get in touch with us at spo@iitk.ac.in.")

		// mail_channel <- mail.GenerateMail("spo@iitk.ac.in",
		// 	"Registration requested on RAS",
		// 	"Company "+pvf.Duration+" has requested to be registered on RAS. The details are as follows:\n\n"+
		// 		"Name: "+pvf.CompanyUniversityName+"\n"+
		// 		"Designation: "+pvf.MentorDesignation+"\n"+
		// 		"Email: "+pvf.MentorEmail+"\n"+
		// 		"Phone: "+pvf.FileName+"\n"+
		// 		"Comments: "+pvf.Duration+"\n")
		// ctx.JSON(http.StatusOK, gin.H{"pid": pid}) // to be removed
		ctx.JSON(http.StatusOK, gin.H{"status": "Successfully Requested"})
	}
}

func deletePVFHandler(ctx *gin.Context) {
	pid, err := util.ParseUint(ctx.Param("pid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = deletePVF(ctx, pid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "deleted PVF"})
}
