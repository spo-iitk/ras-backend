package rc

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/util"
)

type AllResumeResponse struct {
	Name          string       `json:"name"`
	Email         string       `json:"email"`
	Sid           uint         `json:"sid"`
	Rsid          uint         `json:"rsid"`
	Resume        string       `json:"resume"`
	Verified      sql.NullBool `json:"verified"`
	ActionTakenBy string       `json:"action_taken_by"`
	RollNo        string       `json:"roll_no"`
	ResumeType    ResumeType   `json:"resume_type"`
}

func getAllResumesHandler(ctx *gin.Context) {
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var resumes []AllResumeResponse
	err = fetchAllResumes(ctx, rid, &resumes)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resumes)
}

func getResumesHandler(ctx *gin.Context) {
	sid, err := util.ParseUint(ctx.Param("sid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var resumes []StudentRecruitmentCycleResume
	err = fetchStudentResume(ctx, sid, &resumes)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resumes)
}

type putResumeVerifyRequest struct {
	Verified bool `json:"verified"`
}

func putResumeVerifyHandler(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		rsid, err := util.ParseUint(ctx.Param("rsid"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var req putResumeVerifyRequest
		
		err = ctx.BindJSON(&req)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := middleware.GetUserID(ctx)

		ok, studentRCID, err := updateResumeVerify(ctx, rsid, req.Verified, user)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !ok {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "resume not found"})
			return
		}

		var student StudentRecruitmentCycle
		err = FetchStudent(ctx, studentRCID, &student)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		logrus.Infof("%v verified resume with id %d, changed state to %v", user, rsid, req.Verified)

		msg := "Dear " + student.Name + "\n\n"
		msg += "Your resume with id " + ctx.Param("rsid") + " has been "
		if req.Verified {
			msg += "ACCEPTED"
		} else {
			msg += "REJECTED"
		}
		mail_channel <- mail.GenerateMail(student.Email, "Action taken on resume", msg)

		ctx.JSON(http.StatusOK, gin.H{"status": true})
	}
}
