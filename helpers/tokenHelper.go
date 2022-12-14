package helpers

import (
	"context"
	"fmt"
	"time"

	"github.com/KingAnointing/go-gin-jwt-project/configs"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var secret_key = ""
var collections = configs.Collections("jwt-user")

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
			ExpiresAt: &jwt.NumericDate{time.Now().Add(time.Hour * 345)},
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret_key))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secret_key))

	return token, refreshToken, err
}

func UpdateAlltoken(claims, refreshClaims, userId string) {
	var updateObj primitive.D
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	updateObj = append(updateObj, bson.E{"token", claims})
	updateObj = append(updateObj, bson.E{"refresh_token", refreshClaims})
	updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{"updated_at", updated_at})

	upsert := true
	filter := bson.M{"user_id": userId}

	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	collections.UpdateOne(ctx, filter, updateObj, &opt)
	return
}

func ValidateToken(userToken string) (claims *SignedDetail, msg string) {
	token, err := jwt.ParseWithClaims(userToken, &SignedDetail{}, func(t *jwt.Token) (interface{}, error) { return []byte(secret_key), nil })
	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*SignedDetail)
	if !ok {
		msg = fmt.Sprintf("The Token is Invalid")
		return
	}

	if claims.VerifyExpiresAt(claims.ExpiresAt.Time, true) {
		msg = fmt.Sprintf("Token has Expired and it is no longer valid")
		return
	}
	return claims,msg
}
