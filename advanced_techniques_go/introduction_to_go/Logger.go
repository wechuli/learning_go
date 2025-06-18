package main

import (
	"log"
	"log/syslog"
)

func Logger() {
	sysLog, err := syslog.New(syslog.LOG_SYSLOG, "systemLog.go")
	if err != nil {
		log.Println(err)
		return
	}
	log.SetOutput(sysLog)
	log.Print("Everything is fine")
}
