consoleAdmin.POST("/{{modelName}}s/", controller.New{{ModelName}}Post)
    consoleAdmin.GET("/{{modelName}}s/", controller.All{{ModelName}}sGet)
    consoleAdmin.GET("/{{modelName}}s/:id", controller.{{ModelName}}ByIDGet)
    consoleAdmin.PUT("/{{modelName}}s/:id", controller.Update{{ModelName}}Put)
    consoleAdmin.DELETE("/{{modelName}}s/:id", controller.{{ModelName}}Delete)
    
    //ph-router-admin don't remove this line