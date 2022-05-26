package mail

import (
	"log"
	"net/smtp"

	"github.com/spo-iitk/ras-backend/config"
	"github.com/spo-iitk/ras-backend/models"
)

type StudentMailingEntity struct {
	Student models.Student
	Msg     []byte
}

type CompanyMailingEntity struct {
	Company models.CompanyHR
	Msg     []byte
}

type Mail interface {
	StudentMailerService(student_mail_channel chan models.Student)
	CompanyMailerService(company_mail_channel chan models.CompanyHR)
}

func StudentMailerService(student_mail_channel chan StudentMailingEntity) {
	mailCounter := 0
	auth := smtp.PlainAuth("", config.EmailUser, config.EmailPass, config.EmailHost)

	for u := range student_mail_channel {
		log.Println("Setting up smtp")

		var to []string

		if config.EmailType == config.IITKMail {
			to = []string{u.Student.IITKEmail}
		} else {
			to = []string{u.Student.PersonalEmail}
		}

		err := smtp.SendMail(config.EmailHost+":"+config.EmailPort, auth, config.EmailUser, to, u.Msg)

		if err != nil {
			log.Printf("ERROR: while mailing user %v, ID %v\n", to, u.Student.ID)
			log.Println(err)
		} else {
			mailCounter += 1
			log.Printf("Mailed %v\n", u.Student.ID)
			log.Printf("Mails sent since inception: %v\n", mailCounter)
		}
	}
}

func CompanyMailerService(company_mail_channel chan CompanyMailingEntity) {
	mailCounter := 0
	auth := smtp.PlainAuth("", config.EmailUser, config.EmailPass, config.EmailHost)

	for u := range company_mail_channel {
		log.Println("Setting up smtp")

		to := []string{u.Company.Email}

		err := smtp.SendMail(config.EmailHost+":"+config.EmailPort, auth, config.EmailUser, to, u.Msg)

		if err != nil {
			log.Printf("ERROR: while mailing user %v, ID %v\n", to, u.Company.ID)
			log.Println(err)
		} else {
			mailCounter += 1
			log.Printf("Mailed %v\n", u.Company.ID)
			log.Printf("Mails sent since inception: %v\n", mailCounter)
		}
	}
}
