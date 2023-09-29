package plugins

import "github.com/spo-iitk/ras-backend/mail"

func NewNoticeNotification(mail_channel chan mail.Mail, id uint, recruitmentCycleID uint, title string, description string, createdBy string) {
	if recruitmentCycleID != 6 {
		return
	}
	var emails []string
	message := "A new notice has been created by " + createdBy + " with title " + title + " in Placement 2023-24 Phase 1.\n\nDescription: " + description + "\n\nClick here to view the notice: https://placement.iitk.ac.in/student/rc/6/notices"
	mail_channel <- mail.GenerateMails(emails, "God Notice: "+title, message)
}
