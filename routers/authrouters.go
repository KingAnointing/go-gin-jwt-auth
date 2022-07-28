package routers

import (
	"github.com/KingAnointing/go-gin-jwt-project/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.GET("test/greeter1", controllers.Greeter1())
	router.POST("user/signup", controllers.SignUp())
}
