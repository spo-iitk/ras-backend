package application

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
		pvf.IsApproved.Valid = true
		pvf.IsApproved.Bool = true

		var jwtExpiration = viper.GetInt("PVF.EXPIRATION")
		pvf.PVFExpiry = sql.NullTime{
			Time:  time.Now().Add(time.Duration(jwtExpiration) * time.Minute),
			Valid: true,
		}

		err = updatePVF(ctx, &pvf)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// token, err := middleware.GeneratePVFToken("akshat23@iitk.ac.in", pid, rid) // for testing
		token, err := middleware.GeneratePVFToken(pvf.MentorEmail, pid, rid)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// logrus.Infof("A Token %s with  id %d", token, pid) // to be removed
		message := pvf.MentorDesignation + " " + pvf.MentorName + ",\n\n" +
			pvf.Name + " (email: " + pvf.IITKEmail + ") has requested your verification for a project/internship they completed under your guidance.\n\n" +
			"To verify the details and electronically sign the Project Verification Form (PVF), please click the link below (valid upto 3days):\n\n\n" +
			"https://placement.iitk.ac.in/verify?token=" + token + "&rcid=" + util.ParseString(rid) + "\n\n" +
			"Your prompt response is appreciated to ensure timely processing of " + pvf.Name + "'s placement applications.\n\n" +
			"Please note:\n" +
			"The PVF verifies the student's involvement and contributions to the project/internship. \n" +
			"Only projects/internships conducted with IIT Kanpur faculty or external organizations require verification. \n" +
			"If you have any questions regarding the PVF process, please don't hesitate to contact the Students' Placement Office at spo@iitk.ac.in.\n\n" +
			"Thank you for your time and support."

		mail_channel <- mail.GenerateMail(pvf.MentorEmail,
			"Project Verification Required for "+pvf.Name+"'s Internship/Project",
			message,
		)

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

func sendVerificationLinkForStudentAllPvfHandler(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sid, err := util.ParseUint(ctx.Param("sid"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		rid, err := util.ParseUint(ctx.Param("rid"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var pvfs []PVF

		err = fetchAllUnverifiedPvfForStudent(ctx, sid, rid, &pvfs)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for _, pvf := range pvfs {
			pvf.IsApproved.Valid = true
			pvf.IsApproved.Bool = true

			var jwtExpiration = viper.GetInt("PVF.EXPIRATION")
			pvf.PVFExpiry = sql.NullTime{
				Time:  time.Now().Add(time.Duration(jwtExpiration) * time.Minute),
				Valid: true,
			}

			err = updatePVF(ctx, &pvf)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			token, err := middleware.GeneratePVFToken(pvf.MentorEmail, pvf.ID, rid)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			message := "Dear " + pvf.MentorName + ",\n\n" +
				pvf.Name + " (email: " + pvf.IITKEmail + ") has requested your verification for a project/internship they completed under your guidance.\n\n" +
				"To verify the details and electronically sign the Project Verification Form (PVF), please click the link below (valid upto 3 days):\n\n\n" +
				"https://placement.iitk.ac.in/verify?token=" + token + "&rcid=" + util.ParseString(rid) + "\n\n" +
				"Your prompt response is appreciated to ensure timely processing of " + pvf.Name + "'s placement applications.\n\n" +
				"Please note:\n" +
				"The PVF verifies the student's involvement and contributions to the project/internship. " +
				"Only projects/internships conducted with IIT Kanpur faculty or external organizations require verification. " +
				"If you have any questions regarding the PVF process, please don't hesitate to contact the Students' Placement Office at spo@iitk.ac.in.\n\n" +
				"Thank you for your time and support."

			mail_channel <- mail.GenerateMail(pvf.MentorEmail,
				"Project Verification Required for "+pvf.Name+"'s Internship/Project",
				message,
			)
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "Successfully Requested for All PVFs"})
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

func putPVFHandlerForAdmin(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pvf PVF

		// rid, err := util.ParseUint(ctx.Param("rid"))
		// if err != nil {
		// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		err := ctx.ShouldBindJSON(&pvf)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if pvf.ID == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "id is required"})
			return
		}

		var oldPvf PVF

		// err = fetchPvfForAdmin(ctx, rid, oldPvf.ID, &oldPvf)
		err = fetchPVF(ctx, pvf.ID, &oldPvf)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = updatePVF(ctx, &pvf)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		message := "Dear " + oldPvf.Name + ",\n\n" +

			"Action has been taken on your Project Verification Form. Kindly check the status on the RAS Portal." +
			"\n\n" +
			"Regards,\n" +
			"Students' Placement Team, IIT kanpur"
		// logrus.Infof("EmaIL : %s OR %s", pvf.IITKEmail, oldPvf.IITKEmail)
		mail_channel <- mail.GenerateMail(oldPvf.IITKEmail,
			"Project Verification Update for "+oldPvf.Name+"'s Internship/Project",
			message,
		)
		ctx.JSON(http.StatusOK, gin.H{"status": "Updated PVF with id " + util.ParseString(pvf.ID)})
	}
}

func getAllStudentPvfHandler(ctx *gin.Context) {
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sid, err := util.ParseUint(ctx.Param("sid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var pvfs []PVF
	err = fetchAllPvfForStudent(ctx, sid, rid, &pvfs)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, pvfs)

}
