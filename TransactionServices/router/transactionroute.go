package router

import (
	"transactionservices/controllers"

	"github.com/gin-gonic/gin"
)

func TransactionStorageRoute(router *gin.Engine) {
	// api to store the transaction
	api := router.Group("/api/transactionservices")
	{
		api.POST("/deposit", controllers.Deposit())         //insert a deposit transaction in deposit database
		api.POST("/withdrwal", controllers.GetAllDeposit()) //insert a withdraw transaction in withdraw database
		api.GET("/deposits", controllers.GetAllDeposit())
		api.GET("/withdrawal", controllers.GetAllWithdrawal())
		//api.GET("/deposits/:accountnumber", controllers.AccountDeposits())
		//api.GET("/withdrawal/:accountnumber", controllers.AccountWithdrawal())

	}
	// api to make a desposit
	router.GET("http://localhost:3000/api/accountservices/account/:accountnumber/solde")
	// api to make a withdrawal
}
