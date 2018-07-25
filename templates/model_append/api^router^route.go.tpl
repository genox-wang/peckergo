consoleAdmin.POST("/{{model_name}}s/", controller.New{{ModelName}}Post)
    consoleAdmin.GET("/{{model_name}}s/", controller.All{{ModelName}}sGet)
    consoleAdmin.GET("/{{model_name}}s/:id", controller.{{ModelName}}ByIDGet)
    consoleAdmin.PUT("/{{model_name}}s/:id", controller.Update{{ModelName}}Put)
    consoleAdmin.DELETE("/{{model_name}}s/:id", controller.{{ModelName}}Delete)
    
    //ph-router-admin don't remove this line