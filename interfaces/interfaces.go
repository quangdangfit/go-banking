package interfaces

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserID  uint
}

type ResponseAccount struct {
	ID      uint
	Name    string
	Balance int
}

type ResponseUser struct {
	ID       uint
	Username string
	Email    string
	Accounts []ResponseAccount
}

type Validation struct {
	Value string
	Valid string
}

type Register struct {
	Username string
	Email    string
	Password string
}

type Login struct {
	Username string
	Password string
}

type ErrResponse struct {
	Message string
}

type Transaction struct {
	gorm.Model
	From   uint
	To     uint
	Amount int
}

type TransactionBody struct {
	UserId uint
	From   uint
	To     uint
	Amount int
}
