package mail

import (
	"crypto/tls"
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
type SMTPServer struct {
	Host      string
	Port      string
	TLSConfig *tls.Config
}

// ServerName ...
func (s *SMTPServer) ServerName() string {
	return s.Host + ":" + s.Port
}

func (mail *Mail) BuildMessage() []byte {
	header := ""
	header += fmt.Sprintf("From: %s\r\n", viper.GetString("MAIL.USER")+"@iitk.ac.in")
	header += fmt.Sprintf("To: %s\r\n", viper.GetStringSlice("MAIL.WEBTEAM"))

	if len(mail.Cc) > 0 {
		header += fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.Cc, ";"))
	}

	header += fmt.Sprintf("Subject: %s | Recruitment Automation Portal\r\n", mail.Subject)
	header += "\r\n" + mail.Body

	header += "\r\n\r\n Best Recruitment Automation Team \r\n This is an auto-generated email. Please do not reply."

	return []byte(header)
}

func MailerService(mail_channel chan Mail) {
	log.Info("Hello mailer")
	user := viper.GetString("MAIL.USER")
	pass := viper.GetString("MAIL.PASS")
	host := viper.GetString("MAIL.HOST")
	port := viper.GetString("MAIL.PORT")

	smtpServer := SMTPServer{Host: host, Port: port}
	smtpServer.TLSConfig = &tls.Config{InsecureSkipVerify: true, ServerName: smtpServer.Host}

	auth := smtp.PlainAuth("", user, pass, host)
	to := viper.GetStringSlice("MAIL.WEBTEAM")

	for u := range mail_channel {
		mail := Mail{}
		mail.Sender = viper.GetString("MAIL.USER") + "@iitk.ac.in"
		mail.To = to
		mail.Cc = u.Cc
		mail.Bcc = u.Bcc
		mail.Subject = u.Subject
		mail.Body = u.Body

		messageBody := mail.BuildMessage()

		err := smtp.SendMail(smtpServer.ServerName(), auth, mail.Sender, to, messageBody)

		if err != nil {
			log.Info("ERROR: while mailing users\n")
			log.Error(err)
		} else {
			log.Info("Mailed successfully\n")
		}
	}
}
