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

func GenerateAllToken(firstName string, lastName string, email string, uid string, userType string) (string, string, error) {
	claims := &SignedDetail{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Uid:       uid,
		UserType:  userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(24 * time.Hour).Unix(),
		},
	}

	refreshClaims := &SignedDetail{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(184 * time.Hour).Unix(),
		},
	}

}
