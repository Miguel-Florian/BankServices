package routers

import (
	"accountservices/controllers"

	"github.com/gin-gonic/gin"
)

func AccountRoute(router *gin.Engine) {
	//All routes related to book comes here

	api := router.Group("/api/accountservices")
	{
		api.POST("/createaccount", controllers.CreateAccount())
		
		api.GET("/accounts", controllers.GetAccounts())
		api.GET("/account/:accountnumber", controllers.GetAccount())
		api.GET("/account/:accountnumber/solde", controllers.GetSoldeAccount())

		api.PATCH("account/deposit/:accountnumber",controllers.DepositToAccount())
		api.PATCH("account/withdraw/:accountnumber",controllers.WithdrawToAccount())
		

		api.DELETE("/account/:accountnumber", controllers.DeleteAccount())
	}

}
