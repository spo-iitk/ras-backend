package mail

import (
	"fmt"
	"net/smtp"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

type Mail struct {
	Sender  string
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Body    string
}

type MailEntity struct {
	// From the docs Sending "Bcc" messages is accomplished by including an email address in the to parameter but not including it in the msg headers.
	to  []string //NOTE: list of emails seperated by comma
	Msg []byte
}

func (mail *Mail) BuildMessage() []byte {
	header := ""
	header += fmt.Sprintf("From: %s\r\n", mail.Sender)
	if len(mail.To) > 0 {
		header += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	}
	if len(mail.Cc) > 0 {
		header += fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.Cc, ";"))
	}
	if len(mail.Bcc) > 0 {
		header += fmt.Sprintf("Bcc: %s\r\n", strings.Join(mail.Bcc, ";"))
	}

	header += fmt.Sprintf("Subject: %s | Recruitment Automation Portal\r\n", mail.Subject)
	header += "\r\n" + mail.Body

	header += "\r\n\r\n This is an auto-generated email. Please do not reply."

	return []byte(header)
}

func MailerService(mail_channel chan MailEntity) {
	user := viper.GetString("MAIL.USER")
	pass := viper.GetString("MAIL.PASS")
	host := viper.GetString("MAIL.HOST")
	port := viper.GetString("MAIL.PORT")
	address := fmt.Sprintf("%s:%s", host, port)

	auth := smtp.PlainAuth("", user, pass, host)
	to := []string{viper.GetString("MAIL.WEBTEAM")}

	for u := range mail_channel {
		err := smtp.SendMail(address, auth, user, to, u.Msg)
		if err != nil {
			log.Printf("ERROR: while mailing users %v\n", u.to)
			log.Println(err)
		} else {
			log.Printf("Mailed %v\n", u.to)
		}
	}
}
