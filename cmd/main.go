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

var mailQueue chan mail.Mail

func main() {
	var g errgroup.Group

	g.Go(func() error {
		return authServer().ListenAndServe()
	})

	mailQueue = make(chan mail.Mail)

	go mail.Service(mailQueue)

	g.Go(func() error {
		return rasServer(mailQueue).ListenAndServe()
	})

	g.Go(func() error {
		return studentServer().ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
