package transactions

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gitlab.com/quangdangfit/gocommon/utils/logger"
	"go-banking/helpers"
	"net/http"
)

type Service interface {
	CreateTransaction(c *gin.Context)
}

type service struct {
	repo Repository
}

func NewService() Service {
	return &service{repo: NewRepository()}
}

func (s *service) prepareResponse(account *Transaction) map[string]interface{} {
	var accRes TransactionResponse
	data, _ := json.Marshal(account)
	json.Unmarshal(data, &accRes)

	res := map[string]interface{}{
		"account": accRes,
	}

	return res
}
func (s *service) prepareListResponse(accounts *[]Transaction) map[string]interface{} {
	var accRes TransactionResponse
	var accountsRes = []TransactionResponse{}

	for _, acc := range *accounts {
		data, _ := json.Marshal(acc)
		json.Unmarshal(data, &accRes)

		accountsRes = append(accountsRes, accRes)
	}

	res := map[string]interface{}{
		"accounts": accountsRes,
	}

	return res
}

func (s *service) CreateTransaction(c *gin.Context) {
	var reqBody TransactionRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := s.repo.CreateTransaction(reqBody.From, reqBody.To, reqBody.Amount)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, helpers.PrepareResponse(nil, err.Error(), ""))
		return
	}

	res := s.prepareResponse(transaction)
	c.JSON(http.StatusOK, helpers.PrepareResponse(res, "OK", ""))
}

func (s *service) GetTransactionsOfAccount(c *gin.Context) {
	auth := c.GetHeader("Authorization")

	transactions, err := s.repo.GetTransactionsOfAccount(auth)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, helpers.PrepareResponse(nil, err.Error(), ""))
		return
	}
	res := s.prepareListResponse(transactions)
	c.JSON(http.StatusOK, helpers.PrepareResponse(res, "OK", ""))
}
