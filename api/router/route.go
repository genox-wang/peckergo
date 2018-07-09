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

	authorized := router.Group("/console")
	authorized.Use(middleware.JWTAuthRequired())
	{
		authorized.
			Use(middleware.AuthRoleRequired(model.RoleAdmin)).POST("/users/", controller.NewUserPost)
		authorized.
			Use(middleware.AuthRoleRequired(model.RoleAdmin)).GET("/users/all", controller.AllUsersGet)
		authorized.
			Use(middleware.AuthRoleRequired(model.RoleAdmin)).PUT("/users/:id", controller.UpdateUserPut)
		authorized.
			Use(middleware.AuthRoleRequired(model.RoleAdmin)).DELETE("/users/:id", controller.UserDelete)
	}

	// api := router.Group("/api")
	// {

	// }
}
