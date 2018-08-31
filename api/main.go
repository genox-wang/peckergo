package main

import (
	_ "github.com/go-sql-driver/mysql"

	"peckergo/api/config"
	"peckergo/api/model"
	"peckergo/api/router"

	"peckergo/api/utils/log"
)

func main() {
	// 崩溃抓取
	defer log.PanicRecover()
	// 初始化日志
	log.Init()
	model.OpenDB()
	defer model.CloseDB()
	port := config.GetInt("serverPort")
	router.Run(port)
}
