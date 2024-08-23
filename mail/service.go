package mail

import (
	"fmt"
	"net/smtp"
	"strings"
	"text/template"
	"time"

	"github.com/sirupsen/logrus"
)

type Mail struct {
	To      []string
	Subject string
	Body    string
}

func (mail *Mail) BuildMessage() []byte {

	type TemplateData struct {
		To      []string
		Subject string
		Body    string
	}

	var message strings.Builder

	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: Recruitment Automation System IITK<%s>\r\n", sender)
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)

	// If mass mailing, BCC all the users
	if len(mail.To) == 1 {
		msg += fmt.Sprintf("To: %s\r\n\r\n", mail.To[0])
	} else {
		msg += fmt.Sprintf("To: Undisclosed Recipients<%s>\r\n\r\n", webteam)
	}

	message.WriteString(msg)

	bodyWithLineBreaks := strings.ReplaceAll(mail.Body, "\n", "<br>")

	tmpl := template.Must(template.New("Template").Parse(DefaultTemplate))
	err := tmpl.Execute(&message, TemplateData{
		Subject: mail.Subject,
		Body:    bodyWithLineBreaks,
	})
	if err != nil {
		logrus.Errorf("Error executing email template: %v", err)
		return nil
	}

	return []byte(message.String())
}

func batchEmails(to []string, batch int) [][]string {
	var batches [][]string
	for i := 0; i < len(to); i += batch {
		end := i + batch

		if end > len(to) {
			end = len(to)
		}

		batches = append(batches, to[i:end])
	}

	return batches
}

func Service(mailQueue chan Mail) {
	addr := fmt.Sprintf("%s:%s", host, port)
	auth := smtp.PlainAuth("", user, pass, host)

	for mail := range mailQueue {
		message := mail.BuildMessage()
		to := append(mail.To, webteam)
		batches := batchEmails(to, batch)
		for _, emailBatch := range batches {
			if err := smtp.SendMail(addr, auth, sender, emailBatch, message); err != nil {
				logrus.Errorf("Error sending mail: %v", emailBatch)
				logrus.Errorf("Error: %v", err)
			}
			time.Sleep(1 * time.Second)
		}
	}
}
