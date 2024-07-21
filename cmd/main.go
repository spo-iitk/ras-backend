package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/spo-iitk/ras-backend/config"
	"github.com/spo-iitk/ras-backend/mail"
	"golang.org/x/sync/errgroup"
)

const (
	readTimeout  = 5 * time.Second
	writeTimeout = 10 * time.Second
)

func main() {
	var g errgroup.Group
	mail_channel := make(chan mail.Mail)

	gin.SetMode(gin.ReleaseMode)

	go mail.Service(mail_channel)

	g.Go(func() error {
		return authServer(mail_channel).ListenAndServe()
	})

	g.Go(func() error {
		return rasServer(mail_channel).ListenAndServe()
	})

	g.Go(func() error {
		return studentServer(mail_channel).ListenAndServe()
	})

	g.Go(func() error {
		return companyServer().ListenAndServe()
	})

	g.Go(func() error {
		return adminRCServer(mail_channel).ListenAndServe()
	})

	g.Go(func() error {
		return adminApplicationServer(mail_channel).ListenAndServe()
	})

	g.Go(func() error {
		return adminStudentServer(mail_channel).ListenAndServe()
	})

	g.Go(func() error {
		return adminCompanyServer().ListenAndServe()
	})
	g.Go(func() error {
		return verificationServer().ListenAndServe()
	})

	log.Println("Starting Server...")
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
