package main

import (
	_ "github.com/go-sql-driver/mysql"

	"peckergo/api/config"
	"peckergo/api/model"
	"peckergo/api/router"
	"peckergo/api/utils/json"
	"peckergo/api/utils/log"
)

func main() {
	json.InitJSON(json.NewJSONiter())
	initLogLevel()
	model.OpenDB()
	defer model.CloseDB()
	router.Run(config.GetInt("serverPort"))
}

func initLogLevel() {
	log.InitLog(&log.Logrus{}, config.GetInt("log.logLevel"))
	log.Warn(config.GetInt("log.logLevel"))
}
