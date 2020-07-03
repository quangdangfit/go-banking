package interfaces

import (
	"github.com/jinzhu/gorm"
)

type Validation struct {
	Value string
	Valid string
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

type ResponseTransaction struct {
	ID     uint
	From   uint
	To     uint
	Amount int
}
