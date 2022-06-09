package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/rc"
	"github.com/spo-iitk/ras-backend/util"
)

type proformaEmailRequest struct {
	EventID uint   `json:"event_id"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func proformaEmailHandler(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request proformaEmailRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
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
		ctx.JSON(http.StatusOK, gin.H{"success": "email sent"})
	}
}

func postEventReminderHandler(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		eid_string := ctx.Param("eid")
		eid, err := util.ParseUint(eid_string)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var event ProformaEvent
		err = fetchEventByID(ctx, eid, &event)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		studentRCID, err := fetchStudentRCIDByEvents(ctx, eid)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		studentEmails, err := rc.FetchStudentEmailBySRCID(ctx, studentRCID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		mail_channel <- mail.GenerateMails(studentEmails, "Reminder for an Event: "+event.Name, "This is a gentle reminder about a event. Please check the event details at "+"event.URL") //! TODO add url
		ctx.JSON(http.StatusOK, gin.H{"success": "email sent"})
	}
}
