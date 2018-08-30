package main

import (
	_ "github.com/go-sql-driver/mysql"

	"peckergo/api/config"
	"peckergo/api/model"
	"peckergo/api/router"
	"peckergo/api/utils/json"

	"peckergo/api/utils/logutils"
)

func main() {
	// 崩溃抓取
	defer logutils.PanicRecover()
	// 初始化日志
	logutils.Init()
	json.InitJSON(json.NewJSONiter())
	model.OpenDB()
	defer model.CloseDB()
	port := config.GetInt("serverPort")
	router.Run(port)
}
