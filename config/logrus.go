package config

import (
	log "github.com/sirupsen/logrus"
)

// func caller() func(*runtime.Frame) (function string, file string) {
// 	return func(f *runtime.Frame) (function string, file string) {
// 		p, _ := os.Getwd()

// 		return "", fmt.Sprintf("%s:%d", strings.TrimPrefix(f.File, p), f.Line)
// 	}
// }

func logrusConfig() {

	// log.SetFormatter(&log.TextFormatter{
	// 	CallerPrettyfier: caller(),
	// 	FieldMap: log.FieldMap{
	// 		log.FieldKeyFile: "caller",
	// 	},
	// })
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}
