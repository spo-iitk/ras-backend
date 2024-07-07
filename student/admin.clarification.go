package student

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/util"
)

type postClarificationRequest struct {
	Clarification string `json:"clarification" binding:"required"`
}

func postClarificationHandler(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sid, err := util.ParseUint(ctx.Param("sid"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var student Student
		err = getStudentByID(ctx, &student, sid)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var request postClarificationRequest
		err = ctx.ShouldBindJSON(&request)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		mail_channel <- mail.GenerateMail(student.IITKEmail, "Asking Clarification", request.Clarification)
		mail_channel <- mail.GenerateMail(
			middleware.GetUserID(ctx),
			"Clarification Requested from "+student.Name,
			"Dear "+middleware.GetUserID(ctx)+
				"Clarification was requested from "+student.Name+
				"\nSent Mail:\n"+
				request.Clarification)

		ctx.JSON(http.StatusOK, gin.H{"status": "Clarification Mail sent"})
	}
}
