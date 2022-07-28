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
	jwt.RegisteredClaims
}

func GenerateAllToken(firstName string, lastName string, email string, uid string, userType string) (string, string, error) {
	claims := &SignedDetail{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Uid:       uid,
		UserType:  userType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{time.Now().Add(time.Hour * 24)},
			// ExpiresAt: time.Now().Local().Add(24 * time.Hour).Unix(),
		},
	}

	refreshClaims := &SignedDetail{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{time.Now().Add(time.Hour * 184)},
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret_key))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secret_key))

	return token, refreshToken, err
}

