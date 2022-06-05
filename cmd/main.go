package main

import (
	"log"
	"time"

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

	g.Go(func() error {
		return authServer().ListenAndServe()
	})

	mail_channel := make(chan mail.Mail)

	go mail.Service(mail_channel)

	g.Go(func() error {
		return rasServer(mail_channel).ListenAndServe()
	})

	g.Go(func() error {
		return studentServer().ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
