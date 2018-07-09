package router

import (
	"console-template/api/config"
	"fmt"

	"console-template/api/utils/log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

// Run to start a gin server by port.
func Run(port int) {
	if router == nil {
		log.Panic("route is nil")
	}
	router.Run(fmt.Sprintf(":%d", port))
}

func init() {
	gin.SetMode(config.GetString("router.logMode"))
	router = gin.Default()
	if config.GetBool("corsEnable") {
		allowCors()
	}
	// setStaticAsset()
	route()
}

func setStaticAsset() {
	router.Static("/static", config.GetString("html.static"))
}

func allowCors() {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	config.AllowMethods = []string{"GET", "PUT", "PATCH", "POST", "DELETE"}

	router.Use(cors.New(config))
}
