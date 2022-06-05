package mail

import (
	"fmt"
	"net/smtp"

	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

type Mail struct {
	To      []string
	Subject string
	Body    string
}

func (mail *Mail) BuildMessage() []byte {
	message := fmt.Sprintf("From: Recruitment Automation Portal IITK<%s>\r\n", sender)
	message += fmt.Sprintf("Subject: %s | Recruitment Automation Portal\r\n", mail.Subject)

	// If mass mailing, BCC all the users
	if len(mail.To) == 1 {
		message += fmt.Sprintf("To: %s\r\n\r\n", mail.To[0])
	} else {
		message += fmt.Sprintf("To: Undisclosed Recipients<%s>\r\n\r\n", viper.GetStringSlice("MAIL.WEBTEAM"))
	}

	message += mail.Body
	message += "\r\n\r\n Best \r\n Recruitment Automation Team \r\n"
	message += "This is an auto-generated email. Please do not reply."

	return []byte(message)
}

func Service(mailQueue chan Mail) {
	addr := fmt.Sprintf("%s:%s", host, port)

	for mail := range mailQueue {
		message := mail.BuildMessage()

		err := smtp.SendMail(addr, auth, sender, mail.To, message)

		if err != nil {
			logrus.Error(err)
		} else {
			logrus.Infoln("Mailed successfully")
		}
	}
}
