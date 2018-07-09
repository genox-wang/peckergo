package main

import (
	_ "github.com/go-sql-driver/mysql"

	"console-template/api/config"
	"console-template/api/model"
	"console-template/api/router"
	"console-template/api/utils/json"
	"console-template/api/utils/log"
)

func main() {
	json.InitJSON(json.NewJSONiter())
	initLogLevel()
	log.Warnf("DB => %s", config.GetString("db.mysql"))
	model.OpenDB(config.GetString("db.mysql"))
	defer model.CloseDB()
	router.Run(config.GetInt("serverPort"))
}

func initLogLevel() {
	log.InitLog(&log.Logrus{}, config.GetInt("log.logLevel"))
	log.Warn(config.GetInt("log.logLevel"))
}
