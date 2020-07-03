package accounts

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gitlab.com/quangdangfit/gocommon/utils/logger"
	"go-banking/helpers"
	"net/http"
)

type Service interface {
	GetAccountByID(c *gin.Context)
	//GetAccountsByUser(c *gin.Context)
}

type service struct {
	repo Repository
}

func NewService() Service {
	return &service{repo: NewRepository()}
}

func (s *service) prepareResponse(account *Account) map[string]interface{} {
	var userRes AccountResponse
	data, _ := json.Marshal(account)
	json.Unmarshal(data, &userRes)

	res := map[string]interface{}{
		"account": userRes,
	}

	return res
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
