package useraccounts

import (
	"go-banking/helpers"
	"go-banking/interfaces"
)

func getAccount(id uint) *interfaces.Account {
	db := helpers.ConnectDB()
	account := &interfaces.Account{}
	if db.Where("id = ? ", id).First(&account).RecordNotFound() {
		return nil
	}
	defer db.Close()
	return account
}

func updateAccount(id uint, amount int) {
	db := helpers.ConnectDB()
	db.Model(&interfaces.Account{}).Where("id = ?", id).Update("balance", amount)
	defer db.Close()
}
