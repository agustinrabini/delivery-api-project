package main

import (
	"delivery-api-project/config"
	"delivery-api-project/controllers"
	"delivery-api-project/db"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadConfigs()
	db := db.InitSQL()

	r := gin.Default()

	router := controllers.NewRouter(r, db)
	router.MapRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}
}
