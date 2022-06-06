package auth

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
	_ "github.com/spo-iitk/ras-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func openConnection() {
	host := viper.GetString("DATABASE.HOST")
	port := viper.GetString("DATABASE.PORT")
	password := viper.GetString("DATABASE.PASSWORD")

	dbName := viper.GetString("DBNAME.AUTH")
	user := dbName + viper.GetString("DATABASE.USER")

	dsn := "host=" + host + " user=" + user + " password=" + password
	dsn += " dbname=" + dbName + " port=" + port + " sslmode=disable TimeZone=Asia/Kolkata"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to auth database: ", err)
		panic(err)
	}

	db = database

	err = db.AutoMigrate(&User{}, &OTP{}, &CompanySignUpRequest{})
	if err != nil {
		log.Fatal("Failed to migrate auth database: ", err)
		panic(err)
	}

	log.Info("Connected to auth database")
}

func init() {
	openConnection()
	go cleanupOTP()
}
