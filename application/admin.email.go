package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/rc"
)

type proformaEmailRequest struct {
	EventID uint   `json:"event_id"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func proformaEmailHandler(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request proformaEmailRequest

		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		studentRCID, err := fetchStudentRCIDByEvents(ctx, request.EventID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		studentEmails, err := rc.FetchStudentEmailBySRCID(ctx, studentRCID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		mail_channel <- mail.GenerateMails(studentEmails, request.Subject, request.Body)
		ctx.JSON(http.StatusOK, gin.H{"status": "email sent"})
	}
}
