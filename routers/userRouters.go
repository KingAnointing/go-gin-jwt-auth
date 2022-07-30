package routers

import (
	"github.com/KingAnointing/go-gin-jwt-project/controllers"
	"github.com/KingAnointing/go-gin-jwt-project/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.Use(middleware.Authentication())
	router.GET("test/greeter2", controllers.Greeter2())
	router.GET("user/:userId", controllers.GetAUser())
	router.GET("users", controllers.GetUsers())
}
