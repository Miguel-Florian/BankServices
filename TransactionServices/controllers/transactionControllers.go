package controllers

import (
	//"github.com/gin-gonic/gin/internal/json"
	"encoding/json"
	"log"
	"io/ioutil"
	"fmt"
	"context"
	"net/http"
	"time"
	"transactionservices/config"
	"transactionservices/model"
	"transactionservices/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
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

		results, err := withdrawalCollection.Find(ctx, bson.M{})

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
type resdata struct{
	Number_account string            
	Typeaccount   string             
	Amount         int64             
	DateCreation   time.Time         
	DateUpdate     time.Time          
	FirstName      string             
	LastName       string             
	Email          string             
	Adresse        string             
	Phone          string 
}
func MakeDeposit()gin.HandlerFunc{
	return func(c *gin.Context){
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := c.Param("accountnumber")
		defer cancel()
		// recuperes les données du compte passé en parametre
		url := fmt.Sprint("http://localhost:3000/api/accountservices/account/:params",&params)
		res,err := http.Get(url)
		if err != nil{
			c.JSON(http.StatusNotFound, responses.Response{Status: http.StatusNotFound, Message: "Account  doesn't exists", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		resDataJson, err := ioutil.ReadAll(res.Body)
		if err != nil{
			log.Fatal(err)
		}//fin recuperation

		//stocke les données de l'api dans data
		var data resdata
		//fmt.Println(string(resDataJson))
		json.Unmarshal(resDataJson,&data)
		lastAmount := data.Amount 
		fmt.Println(lastAmount)
		//fin et extraction de l'ancien solde

		var depot model.Transaction
		
		if err := c.BindJSON(&depot); err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		//using the library validation to validate required fields
		if validationErr := validate.Struct(&depot); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}
		newDepot := model.Transaction{
			Number_account: 	depot.Number_account,
			Amount:				depot.Amount,
			Date:				time.Now(),
		}
		depotAmout := newDepot.Amount
		result, err := depositCollection.InsertOne(ctx, newDepot)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		url1 := fmt.Sprint("http://localhost:3000/api/accountservices/account/:params/:params2",&params,&depotAmout)
		req,err := http.NewRequest(http.MethodPut,url1,nil)
		if err != nil{
			c.JSON(http.StatusNotFound, responses.Response{Status: http.StatusNotFound, Message: "Account  doesn't exists", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		req.Header.Set("content-Type","application/json")

		client := &http.Client{}
		resp,err := client.Do(req)
		if err != nil{
			log.Fatal(err)
		}
		defer resp.Body.Close()

		c.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "Depot successfully done", Data: map[string]interface{}{"Amount": result}})
		
	}
}
func MakeWithdraw()gin.HandlerFunc{
	return func (c *gin.Context){
		fmt.Println("Retrait")
	}
}