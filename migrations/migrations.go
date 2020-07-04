package migrations

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-banking/api/accounts"
	"go-banking/api/transactions"
	"go-banking/api/users"
	"go-banking/database"
)

func Migrate() {
	database.InitDatabase()
	User := users.User{}
	Account := &accounts.Account{}
	Transactions := &transactions.Transaction{}

	database.DB.AutoMigrate(&User, &Account, &Transactions)
}
