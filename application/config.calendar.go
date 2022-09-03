package application

import (
	"context"
	"log"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

var cal_srv *calendar.Service

func gCalendarConnect() {
	ctxb := context.Background()
	srv, err := calendar.NewService(ctxb, option.WithCredentialsFile("./secrets.GCPcredentials.json"))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}
	cal_srv = srv
}
