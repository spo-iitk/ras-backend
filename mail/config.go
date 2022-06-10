package mail

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "github.com/spo-iitk/ras-backend/config"
)

var (
	user    string
	pass    string
	host    string
	port    string
	webteam string
	// batch   int
	sender string
)

func init() {
	logrus.Info("Initializing mailer")

	user = viper.GetString("MAIL.USER")
	sender = user + "@iitk.ac.in"

	pass = viper.GetString("MAIL.PASS")
	host = viper.GetString("MAIL.HOST")
	port = viper.GetString("MAIL.PORT")
	webteam = viper.GetString("MAIL.WEBTEAM")

	// batch = viper.GetInt("MAIL.BATCH")
}
