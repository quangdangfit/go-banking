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

func (r *repo) generateAccountNumber() string {
	accNumber := strconv.Itoa(int(time.Now().UnixNano()))
	return accNumber
}

// Create function Transaction
//func Transaction(userId uint, from uint, to uint, amount int, jwt string) map[string]interface{} {
//	// Conver uint to string
//	userIdString := fmt.Sprint(userId)
//	// Validate ownership
//	isValid := helpers.ValidateToken(userIdString, jwt)
//	if isValid {
//		// Take sender and receiver
//		fromAccount := getAccount(from)
//		toAccount := getAccount(to)
//		// Handle errors
//		if fromAccount == nil || toAccount == nil {
//			return map[string]interface{}{"message": "Account not found"}
//		} else if fromAccount.UserID != userId {
//			return map[string]interface{}{"message": "You are not owner of the account"}
//		} else if int(fromAccount.Balance) < amount {
//			return map[string]interface{}{"message": "Account balance is too small"}
//		}
//		// Update account
//		updatedAccount := updateAccount(from, int(fromAccount.Balance)-amount)
//		updateAccount(to, int(toAccount.Balance)+amount)
//
//		// Create transaction
//		transactions.CreateTransaction(from, to, amount)
//
//		// Return response
//		var response = map[string]interface{}{"message": "all is fine"}
//		response["data"] = updatedAccount
//		return response
//	} else {
//		return map[string]interface{}{"message": "Not valid token"}
//	}
//}
