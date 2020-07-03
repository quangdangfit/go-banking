package accounts

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gitlab.com/quangdangfit/gocommon/utils/logger"
	"go-banking/helpers"
	"net/http"
)

type Service interface {
	CreateAccount(c *gin.Context)
	GetAccountByID(c *gin.Context)
	GetAccountsByUser(c *gin.Context)
}

type service struct {
	repo Repository
}

func NewService() Service {
	return &service{repo: NewRepository()}
}

func (s *service) prepareResponse(account *Account) map[string]interface{} {
	var accRes AccountResponse
	data, _ := json.Marshal(account)
	json.Unmarshal(data, &accRes)

	res := map[string]interface{}{
		"account": accRes,
	}

	return res
}
func (s *service) prepareListResponse(accounts *[]Account) map[string]interface{} {
	var accRes AccountResponse
	var accountsRes = []AccountResponse{}

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

func (s *service) CreateAccount(c *gin.Context) {
	auth := c.GetHeader("Authorization")

	userUUID, isValid := helpers.ValidateToken(auth)
	if isValid {
		account, err := s.repo.CreateAccount(userUUID, 0)
		if err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, helpers.PrepareResponse(nil, err.Error(), ""))
			return
		}

		res := s.prepareResponse(account)
		c.JSON(http.StatusOK, helpers.PrepareResponse(res, "OK", ""))
	} else {
		c.JSON(http.StatusBadRequest, helpers.PrepareResponse(nil, "token is invalid", ""))
	}
}

func (s *service) GetAccountByID(c *gin.Context) {
	accId := c.Param("uid")
	auth := c.GetHeader("Authorization")

	account, err := s.repo.GetAccount(accId, auth)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, helpers.PrepareResponse(nil, err.Error(), ""))
		return
	}
	res := s.prepareResponse(account)
	c.JSON(http.StatusOK, helpers.PrepareResponse(res, "OK", ""))
}

func (s *service) GetAccountsByUser(c *gin.Context) {
	auth := c.GetHeader("Authorization")

	accounts, err := s.repo.GetAccountsByUser(auth)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, helpers.PrepareResponse(nil, err.Error(), ""))
		return
	}
	res := s.prepareListResponse(accounts)
	c.JSON(http.StatusOK, helpers.PrepareResponse(res, "OK", ""))
}
