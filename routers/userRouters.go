package routers

import (
	"github.com/KingAnointing/go-gin-jwt-project/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("test/greeter2", controllers.Greeter2())
	router.GET("user/:userId", controllers.GetAUser())
}
