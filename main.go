package main

import (
	"github.com/KingAnointing/go-gin-jwt-project/configs"
	"github.com/KingAnointing/go-gin-jwt-project/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.DatabaseConnection()
	port := configs.ConnectionPort()
	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routers.AuthRoutes(router)
	routers.UserRoutes(router)

	router.Run(":" + port)
}
