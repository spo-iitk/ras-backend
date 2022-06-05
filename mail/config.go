package mail

import (
	"net/smtp"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	user string
	pass string
	host string
	port string

	sender string
	auth   smtp.Auth
)

func init() {
	logrus.Info("Initializing mailer")

	user = viper.GetString("MAIL.USER")
	pass = viper.GetString("MAIL.PASS")
	host = viper.GetString("MAIL.HOST")
	port = viper.GetString("MAIL.PORT")

	sender = user + "@iitk.ac.in"
	auth = smtp.PlainAuth("", user, pass, host)
}
