package accounts

import (
	"errors"
	"github.com/google/uuid"
	"go-banking/database"
	"go-banking/helpers"
	"strconv"
	"time"
)

type Repository interface {
	UpdateAccount(uuid string, amount int) (*Account, error)
	GetAccount(uuid string, auth string) (*Account, error)
	CreateAccount(userUUID string, balance uint) (*Account, error)
	GetAccountsByUser(auth string) (*[]Account, error)
	generateAccountNumber() string
}

type repo struct {
}

func NewRepository() Repository {
	return &repo{}
}

func (r *repo) UpdateAccount(uuid string, amount int) (*Account, error) {
	account := Account{}

	database.DB.Where("uuid = ? ", uuid).First(&account)
	account.Balance = uint(amount)
	database.DB.Save(&account)

	return &account, nil
}

// Refactor function getAccount to use database package
func (r *repo) GetAccount(uuid string, auth string) (*Account, error) {
	userUUID, isValid := helpers.ValidateToken(auth)
	if isValid {
		account := Account{}
		if database.DB.Where("uuid = ? ", uuid).First(&account).RecordNotFound() {
			return nil, errors.New("not found account")
		}
		if account.UserUUID == userUUID {
			return &account, nil
		}
	}
	return nil, errors.New("token is invalid")

}

func (r *repo) CreateAccount(userUUID string, balance uint) (*Account, error) {
	account := Account{
		UUID:      uuid.New().String(),
		AccNumber: r.generateAccountNumber(),
		Type:      "standard",
		Balance:   balance,
		UserUUID:  userUUID,
	}
	database.DB.Create(&account)

	return &account, nil
}

func (r *repo) GetAccountsByUser(auth string) (*[]Account, error) {
	userUUID, isValid := helpers.ValidateToken(auth)
	if isValid {
		account := []Account{}
		if database.DB.Where("user_uuid = ? ", userUUID).Find(&account).RecordNotFound() {
			return nil, errors.New("not found account")
		}
		return &account, nil
	}
	return nil, errors.New("token is invalid")
}

func (r *repo) generateAccountNumber() string {
	accNumber := strconv.Itoa(int(time.Now().UnixNano()))
	return accNumber
}
