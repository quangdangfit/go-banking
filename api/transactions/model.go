package transactions

import (
	"github.com/jinzhu/gorm"
)

type Transaction struct {
	gorm.Model
	UUID   string
	From   string
	To     string
	Amount uint
}

type TransactionRequest struct {
	From   string
	To     string
	Amount uint
}

type TransactionResponse struct {
	UUID   string
	From   string
	To     string
	Amount uint
}
