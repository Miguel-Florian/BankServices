package controllers

import (
	"context"
	"net/http"
	"time"
	"transactionservices/config"
	"transactionservices/model"
	"transactionservices/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var depositCollection *mongo.Collection = config.GetCollection(config.DB, "deposits")
var withdrawalCollection *mongo.Collection = config.GetCollection(config.DB, "withdrawals")
var validate = validator.New()

func GetAllDeposit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var deposits []model.Transaction
		defer cancel()

		results, err := depositCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer results.Close(ctx)
		for results.Next(ctx) {
			var deposit model.Transaction
			if err = results.Decode(&deposit); err != nil {
				c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			deposits = append(deposits, deposit)
		}
		c.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": deposits}})
	}
}
func GetAllWithdraws() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var withdrawals []model.Transaction
		defer cancel()

		results, err := accountCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer results.Close(ctx)
		for results.Next(ctx) {
			var withdrawal model.Transaction
			if err = results.Decode(&withdrawal); err != nil {
				c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			withdrawals = append(withdrawals, withdrawal)
		}
		c.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": withdrawals}})
	}
}
func Deposit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var withdraw model.Transaction
		defer cancel()
		// validating request body
		if err := c.BindJSON(&withdraw); err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		//using the library validation to validate required fields
		if validationErr := validate.Struct(&withdraw); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}
		newwithdraw := model.Transaction{
			ID:             primitive.NewObjectID(),
			Number_account: withdraw.Number_account,
			Amount:         withdraw.Amount,
			Date:           time.Now(),
		}
		result, err := depositCollection.InsertOne(ctx, newwithdraw)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusCreated, responses.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}
func Withdraw() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var withdraw model.Transaction
		defer cancel()
		// validating request body
		if err := c.BindJSON(&withdraw); err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		//using the library validation to validate required fields
		if validationErr := validate.Struct(&withdraw); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}
		newwithdraw := model.Transaction{
			ID:             primitive.NewObjectID(),
			Number_account: withdraw.Number_account,
			Amount:         withdraw.Amount,
			Date:           time.Now(),
		}
		result, err := depositCollection.InsertOne(ctx, newwithdraw)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusCreated, responses.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}
