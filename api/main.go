package main

import (
	_ "github.com/go-sql-driver/mysql"

	"peckergo/api/config"
	"peckergo/api/model"
	"peckergo/api/router"
	"peckergo/api/utils/json"

	log "github.com/sirupsen/logrus"
)

func main() {
	initLog()
	json.InitJSON(json.NewJSONiter())
	model.OpenDB()
	defer model.CloseDB()
	port := config.GetInt("serverPort")
	router.Run(port)
}

func initLog() {
	level := log.Level(config.GetInt("log.logLevel"))
	log.SetLevel(level)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.Warnf("logLevel: [%+v]", level)
}
