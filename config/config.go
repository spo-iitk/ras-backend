package config

import (
	"os"
)

const (
	IITKMail     = "IITKMail"
	PersonalMail = "PersonalMail"
)

var (
	EmailUser = os.Getenv("EMAIL_USER")
	EmailPass = os.Getenv("EMAIL_PASS")
	EmailHost = os.Getenv("EMAIL_HOST")
	EmailPort = os.Getenv("EMAIL_PORT")
	EmailType = os.Getenv("EMAIL_TYPE") // IITK or Personal
)

func CfgInit() {}
