package main

import (
	_ "github.com/go-sql-driver/mysql"

	_ "peckergo/api/task"

	"peckergo/api/config"
	"peckergo/api/model"
	"peckergo/api/mq"
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

	mq.Init()
	port := config.GetInt("serverPort")

	router.Run(port)
}
