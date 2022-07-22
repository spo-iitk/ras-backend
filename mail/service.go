package mail

import (
	"fmt"
	"net/smtp"

	"github.com/sirupsen/logrus"
)

type Mail struct {
	To      []string
	Subject string
	Body    string
}

func (mail *Mail) BuildMessage() []byte {
	message := fmt.Sprintf("From: Recruitment Automation System IITK<%s>\r\n", sender)
	message += fmt.Sprintf("Subject: %s | Recruitment Automation System\r\n", mail.Subject)

	// If mass mailing, BCC all the users
	if len(mail.To) == 1 {
		message += fmt.Sprintf("To: %s\r\n\r\n", mail.To[0])
	} else {
		message += fmt.Sprintf("To: Undisclosed Recipients<%s>\r\n\r\n", webteam)
	}

	message += mail.Body
	message += "\r\n\r\n--\r\nRecruitment Automation Sysytem\r\n"
	message += "Indian Institute of Technology Kanpur\r\n\r\n"
	message += "This is an auto-generated email. Please do not reply."

	return []byte(message)
}

func Service(mailQueue chan Mail) {
	addr := fmt.Sprintf("%s:%s", host, port)
	auth := smtp.PlainAuth("", user, pass, host)

	for mail := range mailQueue {
		message := mail.BuildMessage()

		// for i := 0; i < len(mail.To); i += batch {
		// 	end := i + batch
		// 	if end > len(mail.To) {
		// 		end = len(mail.To)
		// 	}

		to := append(mail.To, webteam)

		if err := smtp.SendMail(addr, auth, sender, to, message); err != nil {
			logrus.Errorf("Error sending mail: %v", to, err)
		}
		// }
	}
}
