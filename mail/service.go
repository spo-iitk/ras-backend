package mail

import (
	"fmt"
	"net/smtp"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type Mail struct {
	To      []string
	Subject string
	Body    string
}

func (mail *Mail) BuildMessage() []byte {
	message := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	message += fmt.Sprintf("From: Recruitment Automation System IITK<%s>\r\n", sender)
	message += fmt.Sprintf("Subject: %s | Recruitment Automation System\r\n", mail.Subject)

	// If mass mailing, BCC all the users
	if len(mail.To) == 1 {
		message += fmt.Sprintf("To: %s\r\n\r\n", mail.To[0])
	} else {
		message += fmt.Sprintf("To: Undisclosed Recipients<%s>\r\n\r\n", webteam)
	}

	message += strings.Replace(mail.Body, "\n", "<br>", -1)
	message += "<br><br>--<br>Recruitment Automation Sysytem<br>"
	message += "Indian Institute of Technology Kanpur<br><br>"
	message += "<small>This is an auto-generated email. Please do not reply.</small>"

	return []byte(message)
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
		batches:= batchEmails(to, batch); 
		for _, emailBatch:= range batches {
			if err := smtp.SendMail(addr, auth, sender, emailBatch, message); err != nil {
				logrus.Errorf("Error sending mail: %v", emailBatch)
				logrus.Errorf("Error: %v", err)
			}
			time.Sleep(1 * time.Second)
		}
	}
}
