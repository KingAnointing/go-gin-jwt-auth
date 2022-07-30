package middleware

import (
	"net/http"

	"github.com/KingAnointing/go-gin-jwt-project/responses"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": "Unauthorized to access this resources"}})
			c.Abort()
			return
		}
	}
}
