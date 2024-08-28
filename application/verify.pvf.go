package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/util"
)

func getPvfForVerificationHandler(ctx *gin.Context) {
	// ctx.JSON(http.StatusOK, gin.H{"pid": middleware.GetPVFID(ctx)})
	pid := middleware.GetPVFID(ctx)

	// pid, err := util.ParseUint(ctx.Param("pid"))
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// rid, err := util.ParseUint(ctx.Param("rid"))
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	rid := middleware.GetRcID(ctx)
	var jps PVF
	err := fetchPvfForVerification(ctx, pid, rid, &jps)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, jps)
}

func putPVFHandler(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pvf PVF

		err := ctx.ShouldBindJSON(&pvf)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		pid := middleware.GetPVFID(ctx)

		pvf.ID = pid

		if pvf.ID == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "id is required"})
			return
		}

		var oldJp PVF
		err = fetchPVF(ctx, pvf.ID, &oldJp)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// jp.ActionTakenBy = middleware.GetUserID(ctx)

		// publishNotice := oldJp.Deadline == 0 && jp.Deadline != 0

		err = updatePVF(ctx, &pvf)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var action string
		if pvf.IsVerified.Bool {
			action = "APPROVED"
		} else {
			action = "DENIED"
		}
		messageMentor := "Dear " + oldJp.MentorName + ",\n\n" +
			"The action " + action + " on the Project Verification Form of the " + oldJp.Name + " has been taken.\n\n" +
			"If you have not done this action please reach out to spo@iitk.ac.in\n\n" +
			"Regards,\n" +
			"Students' Placement Team, IIT kanpur"

		mail_channel <- mail.GenerateMail(oldJp.MentorEmail,
			"Project Verification Update for "+oldJp.Name+"'s Internship/Project",
			messageMentor,
		)
		messageStudent := "Dear " + oldJp.Name + ",\n\n" +

			"Action has been taken on your Project Verification Form. Kindly check the status on the RAS Portal." +
			"\n\n" +
			"Regards,\n" +
			"Students' Placement Team, IIT kanpur"

		mail_channel <- mail.GenerateMail(oldJp.IITKEmail,
			"Project Verification Update for "+oldJp.Name+"'s Internship/Project",
			messageStudent,
		)
		ctx.JSON(http.StatusOK, gin.H{"status": "Updated PVF with id " + util.ParseString(pvf.ID)})

	}
}
