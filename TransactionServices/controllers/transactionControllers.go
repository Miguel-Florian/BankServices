package controllers

import (
	//"github.com/gin-gonic/gin/internal/json"
	"encoding/json"
	"bytes"
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
		url := fmt.Sprintf("http://localhost:3000/api/accountservices/account/%s",params)
		res,err := http.Get(url)
		if err != nil{
			c.JSON(http.StatusNotFound, responses.Response{Status: http.StatusNotFound, Message: "Account  doesn't exists", Data: map[string]interface{}{"data": err.Error()}})
			return
		}else{
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
				ID : 				primitive.NewObjectID(),
				Number_account: 	depot.Number_account,
				Amount:				depot.Amount,
				Date:				time.Now(),
			}
			//fmt.Println(newDepot)

			var depotJson model.Transaction
			nd,_ := json.Marshal(newDepot)
			er := json.Unmarshal(nd, &depotJson)
			if er != nil {
    			fmt.Println(er)
			}
			//fmt.Println(depotJson)
			depotAmount := depotJson.Amount
			fmt.Println(depotAmount)
			result, err := depositCollection.InsertOne(ctx, newDepot)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		
			payload := map[string]int64{
				"amount": depotAmount + lastAmount,
			}
			// encode payload to JSON
			jsonPayload, err := json.Marshal(payload)
			if err != nil {
				panic(err)
			}
			fmt.Println(payload)
			url1 := fmt.Sprintf("http://localhost:3000/api/accountservices/account/%s",params)
			fmt.Println(url1)
			req,err := http.NewRequest("PATCH",url1,bytes.NewBuffer(jsonPayload))
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
}

func MakeWithdraw()gin.HandlerFunc{
	return func (c *gin.Context){
		fmt.Println("Retrait")
	}
}