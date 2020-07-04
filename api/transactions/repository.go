package transactions

import (
	"errors"
	"gitlab.com/quangdangfit/gocommon/utils/logger"
	"go-banking/api/accounts"
	"go-banking/database"
)

type Repository interface {
	CreateTransaction(From string, To string, Amount uint) (*Transaction, error)
	GetTransactionsByID(uuid string) (*Transaction, error)
	GetTransactionsOfAccount(accNumber string) (*[]Transaction, error)
}

type repo struct {
}

func NewRepository() Repository {
	return &repo{}
}

func (r *repo) CreateTransaction(From string, To string, Amount uint) (*Transaction, error) {
	accRepo := accounts.NewRepository()
	fromAcc, err := accRepo.GetAccount(From, "")
	if err != nil {
		logger.Error("Not found from account")
		return nil, err
	}

	if fromAcc.Balance < Amount {
		return nil, errors.New("balance is not enough")
	}

	toAcc, err := accRepo.GetAccount(From, "")
	if err != nil {
		logger.Error("Not found to account")
		return nil, err
	}

	accRepo.UpdateAccount(fromAcc.UUID, int(fromAcc.Balance-Amount))
	accRepo.UpdateAccount(toAcc.UUID, int(toAcc.Balance+Amount))

	transaction := &Transaction{From: From, To: To, Amount: Amount}
	database.DB.Create(&transaction)

	return transaction, nil
}

func (r *repo) GetTransactionsByID(uuid string) (*Transaction, error) {
	var transaction Transaction
	if database.DB.Where("uuid = ? ", uuid).First(&transaction).RecordNotFound() {
		return nil, errors.New("user not found")
	}
	return &transaction, nil
}

func (r *repo) GetTransactionsOfAccount(accNumber string) (*[]Transaction, error) {
	// Find and return transactions
	var transactions []Transaction
	database.DB.Where("from = ?", accNumber).Or("to = ?", accNumber).Find(&transactions)

	return &transactions, nil
}
