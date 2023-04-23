package controllers

import (
	"accountservices/config"
	"accountservices/models"
	"accountservices/responses"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// create a account table in the database accountservices
var accountCollection *mongo.Collection = config.GetCollection(config.DB, "accounts")
var validate = validator.New()

//const SecretKey = "SECRETKEY"

// create an account
func CreateAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var account models.Account
		defer cancel()
		// validatind request body
		if err := c.BindJSON(&account); err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		//using the library validation to validate required fields
		if validationErr := validate.Struct(&account); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}
		pass, _ := bcrypt.GenerateFromPassword([]byte(account.Password), 14)
		password := string(pass[:])
		newAccount := models.Account{
			ID:             primitive.NewObjectID(),
			Number_account: account.Number_account,
			Type_account:   account.Type_account,
			Amount:         account.Amount,
			DateCreation:   time.Now(),
			DateUpdate:     time.Now(),
			FirstName:      account.FirstName,
			LastName:       account.LastName,
			Email:          account.Email,
			Adresse:        account.Adresse,
			Phone:          account.Phone,
			Password:       password,
		}
		result, err := accountCollection.InsertOne(ctx, newAccount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusCreated, responses.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetAccounts() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var accounts []models.Account
		defer cancel()

		results, err := accountCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleaccount models.Account
			if err = results.Decode(&singleaccount); err != nil {
				c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			accounts = append(accounts, singleaccount)
		}
		c.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": accounts}})
	}
}

func GetAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := c.Param("accountnumber")
		var account models.Account
		defer cancel()

		filter := bson.D{{
			Key:   "account_number",
			Value: params,
		}}
		err := accountCollection.FindOne(ctx, filter).Decode(&account)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": account}})
	}
}

func GetSoldeAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := c.Param("accountnumber")
		var account models.Account
		defer cancel()

		filter := bson.D{{
			Key:   "account_number",
			Value: params,
		}}
		err := accountCollection.FindOne(ctx, filter).Decode(&account)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": account.Amount}})
	}
}

func DeleteAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := c.Param("accountnumber")
		defer cancel()

		filter := bson.D{{
			Key:   "account_number",
			Value: params,
		}}
		result, err := accountCollection.DeleteOne(ctx, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.Response{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Book with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Successfully deleted!"}},
		)
	}
}
