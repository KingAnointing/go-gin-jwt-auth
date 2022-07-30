package middleware

import (
	"net/http"

	"github.com/KingAnointing/go-gin-jwt-project/helpers"
	"github.com/KingAnointing/go-gin-jwt-project/responses"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": "Authorization token no provided"}})
			c.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err}})
			c.Abort()
			return
		}

		c.Set("first_name",claims.FirstName)
		c.Set("last_name",claims.LastName)
		c.Set("email",claims.Email,)
		c.Set("uid",claims.Uid)
		c.Set("user_type",claims.UserType)
		c.Next()
	}
}
