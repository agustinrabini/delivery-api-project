package main

import (
	"delivery-api-project/config"
	"delivery-api-project/controllers"
	"delivery-api-project/db"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title Delivery Project <Domain>
// @version 1.0
// @description This is a project of an api.
// @termsOfService http://swagger.io/terms/

// @contact.name Agustin Rabini
// @contact.url http://www.swagger.io/support
// @contact.email agustinrabini99@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {

	config.LoadConfigs()
	db := db.InitSQL()

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//r.Use(middleware.Middleware99())

	router := controllers.NewRouter(r, db)
	router.MapRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}
}
