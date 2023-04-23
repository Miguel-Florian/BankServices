package main

import (
	"accountservices/config"
	"accountservices/routers"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting server ...")
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

	routers.AccountRoute(r)

	r.SetTrustedProxies(nil)
	r.Run(PORT)
}
