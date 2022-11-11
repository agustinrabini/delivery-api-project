package main

import (
	"delivery-api-project/config"
	"delivery-api-project/controllers"
	"delivery-api-project/db"

	"github.com/gin-gonic/gin"
)

func main() {

	db := db.InitSQL()
	config.LoadConfigs()

	r := gin.Default()

	router := controllers.NewRouter(r, db)
	router.MapRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}
}
