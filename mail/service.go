package mail

import (
	"fmt"
	"net/smtp"
	"strings"
	"time"
	"text/template" 
	"github.com/sirupsen/logrus"
	"crypto/tls"
)

type Mail struct {
	To      []string
	Subject string
	Body    string
}

func (mail *Mail) BuildMessage() []byte {

	type TemplateData struct {
		To      []string
		Subject string
		Body    string
	}

	var message strings.Builder

	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: Recruitment Automation System IITK<%s>\r\n", sender)
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)

	// If mass mailing, BCC all the users
	if len(mail.To) == 1 {
		msg += fmt.Sprintf("To: %s\r\n\r\n", mail.To[0])
	} else {
		msg += fmt.Sprintf("To: Undisclosed Recipients<%s>\r\n\r\n", webteam)
	}

	message.WriteString(msg)

	bodyWithLineBreaks := strings.ReplaceAll(mail.Body, "\n", "<br>")

	tmpl := template.Must(template.New("Template").Parse(DefaultTemplate))
	err := tmpl.Execute(&message, TemplateData{
		Subject: mail.Subject,
		Body:    bodyWithLineBreaks,
	})
	if err != nil {
		logrus.Errorf("Error executing email template: %v", err)
		return nil
	}

	return []byte(message.String())
}

func batchEmails(to []string, batch int) [][]string {
	var batches [][]string
	for i := 0; i < len(to); i += batch {
		end := i + batch

		if end > len(to) {
			end = len(to)
		}

		batches = append(batches, to[i:end])
	}

	return batches
}

func Service(mailQueue chan Mail) {
	addr := fmt.Sprintf("%s:%s", host, port)
	auth := smtp.PlainAuth("", user, pass, host)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // Change this to verify server certificate in production
	}
	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		logrus.Errorf("Error dialing SMTPS server: %v", err)
	}

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		logrus.Errorf("Error creating SMTP client: %v", err)
	}


	if err := client.Auth(auth); err != nil {
		logrus.Errorf("Error authenticating: %v", err)
	}

	for mail := range mailQueue {
		message := mail.BuildMessage()
		to := append(mail.To, webteam)
		batches := batchEmails(to, batch)
		for _, emailBatch := range batches {
			
			// if err := client.StartTLS(tlsConfig); err != nil {
			// 	logrus.Errorf("Error starting TLS: %v", err)
			// 	continue
			// }
			if err := client.Mail(sender); err != nil {
				logrus.Errorf("Error setting sender: %v", err)
			}

			for _, recipient := range emailBatch {
				if err := client.Rcpt(recipient); err != nil {
					logrus.Errorf("Error setting recipient %s: %v", recipient, err)
					continue
				}
			}

			// Send the email body
			w, err := client.Data()
			if err != nil {
				logrus.Errorf("Error creating data writer: %v", err)
				continue
			}
			_, err = w.Write(message)
			if err != nil {
				logrus.Errorf("Error writing email body: %v", err)
			}
			err = w.Close()
			if err != nil {
				logrus.Errorf("Error closing data writer: %v", err)
			}
			time.Sleep(1 * time.Second)
		}
	}

	if err := client.Quit(); err != nil {
		logrus.Errorf("Error closing connection: %v", err)
	}
}
