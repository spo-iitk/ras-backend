package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type alertMsg struct {
	Endpoint string      `json:"endpoint"`
	Err      interface{} `json:"error"`
}

var alertChannel chan alertMsg

var unix_socket = "/tmp/ras-backend.sock"

func sendAlertToDiscord() {
	conn, err := net.Dial("unix", unix_socket)
	for err != nil {
		// logrus.Error("Error in connecting to socket: ", err)
		conn, err = net.Dial("unix", unix_socket)
		time.Sleep(5 * time.Second)
	}
	defer conn.Close()
	log.Println("Ready to send panic alerts")
	for {
		alert := <-alertChannel
		jsonData, err := json.Marshal(alert)
		if err != nil {
			logrus.Error("Error in alerting panic: ", err)
			continue
		}
		_, err = conn.Write(jsonData)
		if err != nil {
			logrus.Error("Error in writing data to socket: ", err)
		}
	}
}

func recoveryHandler(c *gin.Context, err interface{}) {
	alertChannel <- alertMsg{c.Request.URL.Path, err}
	c.AbortWithStatus(http.StatusInternalServerError)
}

func init() {
	alertChannel = make(chan alertMsg)
	go sendAlertToDiscord()
}
