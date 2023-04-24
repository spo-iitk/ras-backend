package config

import (
	"github.com/sirupsen/logrus"
)

func logrusConfig() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)

	// f, err := os.OpenFile("raslog.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	// if err != nil {
	// 	fmt.Printf("error opening file: %v", err)
	// 	panic(err)
	// }
	// logrus.SetOutput(f)
}
