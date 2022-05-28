package auth

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func OpenConnection() {
	host := viper.GetString("DATABASE.HOST")
	port := viper.GetString("DATABASE.PORT")
	user := viper.GetString("AUTH.DATABASE.USER")
	password := viper.GetString("AUTH.DATABASE.PASSWORD")
	dbName := viper.GetString("AUTH.DATABASE.NAME")

	dsn := "host=" + host + " user=" + user + " password=" + password
	dsn += " dbname=" + dbName + " port=" + port + " sslmode=disable TimeZone=Asia/Kolkata"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	db = database
	db.AutoMigrate(&User{}, &Role{}, &OTP{})

	log.Info("Connected to database")
}
