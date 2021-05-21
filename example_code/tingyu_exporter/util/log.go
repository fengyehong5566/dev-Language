package util

import (
	"github.com/prometheus/common/log"
	"os"
)

var Logger log.Logger
func CreateLogger() log.Logger {
	conf := ReadConf()
	logf,err := os.OpenFile("./log/exporter.log", os.O_APPEND|os.O_CREATE|os.O_RDWR|os.O_TRUNC,0644)
	if err != nil {
		log.Fatalf("Open %s fail. err: %s\n","./log/exporter.log", err)
	}
	Logger = log.NewLogger(logf)
	//Logger.SetFormat()
	err = Logger.SetLevel(conf.LogLevel)
	if err != nil {
		Logger.Fatalln("Failed to initialize the log level")
	}
	return Logger
}
