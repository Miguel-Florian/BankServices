package router

import (
	"transactionservices/controllers"

	"github.com/gin-gonic/gin"
)

func TransactionStorageRoute(router *gin.Engine) {
	// api to store the transaction
	api := router.Group("/api/transactionservices")
	{
		api.POST("/deposit/:accountnumber",controllers.MakeDeposit())
		api.POST("/withdraw/:accountnumber",controllers.MakeWithdraw())

		api.GET("/deposits", controllers.GetAllDeposit())
		api.GET("/withdraws", controllers.GetAllWithdraws())

	}
}
