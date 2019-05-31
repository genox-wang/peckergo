package router

import (

	// "net/http"
	// "peckergo-api/app/config"
	"peckergo/api/controller"
	"peckergo/api/middleware"
	"peckergo/api/model"
	// "github.com/gin-gonic/gin"
)

func route() {
	router.POST("/console/login", controller.LoginPost)
	router.GET("/console/captcha", controller.CaptchaGet)
	// router.LoadHTMLFiles(config.GetString("html.index"))

	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", nil)
	// })

	consoleAdmin := router.Group("/console")
	consoleAdmin.
		Use(middleware.JWTAuthRequired()).
		Use(middleware.AuthRoleRequired(model.RoleAdmin))
	{
		consoleAdmin.POST("/users/", controller.NewUserPost)
		consoleAdmin.GET("/users/", controller.AllUsersGet)
		consoleAdmin.GET("/users/:id", controller.UserByIDGet)
		consoleAdmin.PUT("/users/:id", controller.UpdateUserPut)
		consoleAdmin.DELETE("/users/:id", controller.UserDelete)

		consoleAdmin.GET("/log_managements/", controller.AllLogManagementsGet)

		consoleAdmin.GET("/dbtabless/", controller.AllDbtablessGet)

		//ph-router-admin don't remove this line
	}

	consoleOperator := router.Group("/console")
	consoleOperator.
		Use(middleware.JWTAuthRequired()).
		Use(middleware.AuthRoleRequired(model.RoleAdmin, model.RoleClient))
	{
		consoleOperator.GET("/map/users/", controller.AllUserIDNameMapGet)
		//ph-router-operator don't remove this line
	}

	// api := router.Group("/api")
	// {

	// }
}
