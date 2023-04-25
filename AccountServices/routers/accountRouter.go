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
		api.PUT("account/:accountnumber/:amount",controllers.UpdateSoldeAccount())
		api.GET("/account/:accountnumber", controllers.GetAccount())
		api.GET("/account/:accountnumber/solde", controllers.GetSoldeAccount())
		api.DELETE("/account/:accountnumber", controllers.DeleteAccount())
	}

}
