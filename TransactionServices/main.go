package main

import (
	"fmt"
	"log"
//	"net/http"
	"os"
	"transactionservices/config"
	"transactionservices/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting Transaction server ...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")

	r := gin.Default()

	r.RedirectTrailingSlash = true
	r.RedirectFixedPath = true

	//run database server
	config.ConnectDB()

	router.TransactionStorageRoute(r)

	r.SetTrustedProxies(nil)
	r.Run(PORT)
}
