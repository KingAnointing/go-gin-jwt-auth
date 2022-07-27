package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)
var secret_key = ""
type SignedDetail struct {
	FirstName string
	LastName  string
	Email     string
	Uid       string
	UserType  string
	jwt.StandardClaims
}


