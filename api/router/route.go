package router

import (

	// "net/http"
	// "trans-trip-admin-api/app/config"
	"console-template/api/controller"
	"console-template/api/middleware"
	"console-template/api/model"
	// "github.com/gin-gonic/gin"
)

func route() {
	router.POST("/console/login", controller.LoginPost)
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
		consoleAdmin.GET("/users/all", controller.AllUsersGet)
		consoleAdmin.PUT("/users/:id", controller.UpdateUserPut)
		consoleAdmin.DELETE("/users/:id", controller.UserDelete)
	}

	// consoleOperator := router.Group("/console")
	// consoleOperator.
	// 	Use(middleware.JWTAuthRequired()).
	// 	Use(middleware.AuthRoleRequired(model.RoleAdmin, model.RoleOperator))
	// {

	// }

	// api := router.Group("/api")
	// {

	// }
}
