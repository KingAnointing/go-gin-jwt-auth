package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("user is unauthorized to access user detail")
		return
	}
	return
}

func MatchUserTypeToId(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	if userType == "USER" && userId != uid {
		err = errors.New("user is unauthorized to access user detail")
		return
	}

	err = CheckUserType(c, userType)
	return
}
