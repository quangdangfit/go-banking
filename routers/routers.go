package routers

import (
	"github.com/gin-gonic/gin"
	"go-banking/accounts"
	"go-banking/users"
)

func API(e *gin.Engine) {
	v1 := e.Group("api/v1")
	{
		user := users.NewService()
		v1.POST("/login", user.Login)
		v1.POST("/register", user.Register)
		v1.GET("/users/:uid", user.GetUser)

		account := accounts.NewService()
		v1.GET("/accounts/:uid", account.GetAccountByID)

	}
}
