package mail

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"

	"github.com/spo-iitk/ras-backend/config"
	"github.com/spo-iitk/ras-backend/models"
)

type Mail struct {
	Sender  string
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Body    string
}

type StudentMailingEntity struct {
	Student models.Student
	Msg     []byte
}

type CompanyMailingEntity struct {
	Company models.CompanyHR
	Msg     []byte
}

type MassMailEntity struct {
	// From the docs Sending "Bcc" messages is accomplished by including an email address in the to parameter but not including it in the msg headers.
	to  []string //NOTE: list of emails seperated by comma
	Msg []byte
}

type Mailer interface {
	MassMailerService(mail_channel chan MassMailEntity)
	StudentMailerService(student_mail_channel chan models.Student)
	CompanyMailerService(company_mail_channel chan models.CompanyHR)
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

	header += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	header += "\r\n" + mail.Body

	return []byte(header)
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

func MassMailerService(mail_channel chan MassMailEntity) {
	auth := smtp.PlainAuth("", config.EmailUser, config.EmailPass, config.EmailHost)
	to := []string{config.EmailWebteam}
	for u := range mail_channel {
		err := smtp.SendMail(config.EmailHost+":"+config.EmailPort, auth, config.EmailUser, to, u.Msg)
		if err != nil {
			log.Printf("ERROR: while mailing users %v\n", u.to)
			log.Println(err)
		} else {
			log.Printf("Mailed %v\n", u.to)
		}
	}
}
