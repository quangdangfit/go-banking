package routers

import (
	"github.com/gin-gonic/gin"
	"go-banking/api/accounts"
	"go-banking/api/transactions"
	"go-banking/api/users"
)

func API(e *gin.Engine) {
	v1 := e.Group("api/v1")
	{
		user := users.NewService()
		v1.POST("/login", user.Login)
		v1.POST("/register", user.Register)
		v1.GET("/users/:uid", user.GetUser)

		account := accounts.NewService()
		v1.POST("/accounts", account.CreateAccount)
		v1.GET("/accounts", account.GetAccountsByUser)
		v1.GET("/accounts/:uid", account.GetAccountByID)

		transaction := transactions.NewService()
		v1.GET("/transactions", transaction.CreateTransaction)
	}
}
