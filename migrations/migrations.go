package migrations

import (
	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-banking/database"
	"go-banking/helpers"
	"go-banking/interfaces"
	"go-banking/users"
)

func createAccounts() {
	usrs := &[2]users.User{
		{Username: "quang", Email: "quang@quang.com"},
		{Username: "dang", Email: "dang@dang.com"},
	}

	for i := 0; i < len(usrs); i++ {
		// Correct one way
		generatedPassword := helpers.HashAndSalt([]byte(usrs[i].Username))
		user := &users.User{Username: usrs[i].Username, Email: usrs[i].Email, Password: generatedPassword, UID: uuid.New().String()}
		database.DB.Create(&user)

		//account := &interfaces.Account{Type: "Daily Account", Name: string(usrs[i].Username + "'s" + " account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
		//database.DB.Create(&account)
	}
}

func Migrate() {
	database.InitDatabase()
	User := users.User{}
	Account := &interfaces.Account{}
	Transactions := &interfaces.Transaction{}

	database.DB.AutoMigrate(&User, &Account, &Transactions)

	createAccounts()
}
