package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/KingAnointing/go-gin-jwt-project/configs"
	"github.com/KingAnointing/go-gin-jwt-project/helpers"
	"github.com/KingAnointing/go-gin-jwt-project/models"
	"github.com/KingAnointing/go-gin-jwt-project/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var collections *mongo.Collection = configs.Collections("jwt-user")
var validate = validator.New()

// greeter function to test API-1
func Greeter1() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Hello from API auth router"}})
	}
}

// greeter function to test API-2
func Greeter2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Hello from API user router"}})
	}
}

func HashPassword(password string) string {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(byte)
}

func VerifyPassword(hashedPassword, password string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPassword))
	isValid := true
	msg := ""
	if err != nil {
		isValid = false
		msg = fmt.Sprintf("Password is Incorrect,%v", err.Error())
		return isValid, msg
	}
	return isValid, msg
}

// signup function --> Creating user and essential item needed
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
			return
		}

		if validateErr := validate.Struct(&user); validateErr != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": validateErr.Error()}})
			return
		}

		count, err := collections.CountDocuments(ctx, bson.M{"email": user.Email})
		if count > 0 {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"error": "Email already Exist in Database"}})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
			return
		}

		count, err = collections.CountDocuments(ctx, bson.M{"phone": user.Phone})
		if count > 0 {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"error": "Phone Number already Exist in Database"}})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
			return
		}

		user.ID = primitive.NewObjectID()
		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UserId = user.ID.Hex()
		token, refreshToken, _ := helpers.GenerateAllToken(*user.FirstName, *user.LastName, *user.Email, user.UserId, *user.UserType)
		user.Token = &token
		user.RefreshToken = &refreshToken

		hashedPassword := HashPassword(*user.Password)
		user.Password = &hashedPassword

		result, err := collections.InsertOne(ctx, &user)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

// login function --> Post login cred like email and password to get logged in
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		var foundUser models.User

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
			return
		}

		if err := collections.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser); err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"error": "Email or Password Incorrect !!"}})
			return
		}

		passwordIsValid, err := VerifyPassword(*user.Password, *foundUser.Password)
		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"error": err}})
			return
		}

		if *foundUser.Email == "" {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"error": "User Not Found !!"}})
			return
		}

		token, refreshtoken, _ := helpers.GenerateAllToken(*foundUser.FirstName, *foundUser.LastName, *foundUser.Email, foundUser.UserId, *foundUser.UserType)

		helpers.UpdateAlltoken(token, refreshtoken, foundUser.UserId)

		if err := collections.FindOne(ctx, bson.M{"user_id": foundUser.UserId}).Decode(&foundUser); err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": foundUser}})
	}
}

// get a single user --> only admin can get all user and regular user can only get thier own profile
func GetAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user models.User
		userId := c.Param("userId")
		if err := helpers.MatchUserTypeToId(c, userId); err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if err := collections.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user); err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
	}
}
