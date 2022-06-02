package config

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func logrusConfig() {
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	f, err := os.OpenFile("raslog.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		panic(err)
	}
	log.SetOutput(f)
}
