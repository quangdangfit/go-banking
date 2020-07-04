package users

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gitlab.com/quangdangfit/gocommon/utils/logger"
	"go-banking/utils/response"
	"net/http"
)

type Service interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	GetUser(c *gin.Context)
}

type service struct {
	repo Repository
}

func NewService() Service {
	return &service{repo: NewRepository()}
}

func (s *service) prepareResponse(user *User, token bool) map[string]interface{} {
	var userRes UserResponse
	data, _ := json.Marshal(user)
	json.Unmarshal(data, &userRes)

	res := map[string]interface{}{
		"user": userRes,
	}
	if token {
		res["token"] = s.repo.GenerateToken(user)
	}

	return res
}

func (s *service) Login(c *gin.Context) {
	var reqBody LoginRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.repo.Login(reqBody.Username, reqBody.Password)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.PrepareResponse(nil, err.Error(), ""))
		return
	}

	res := s.prepareResponse(user, true)
	c.JSON(http.StatusOK, response.PrepareResponse(res, "OK", ""))
}

func (s *service) Register(c *gin.Context) {
	var reqBody RegisterRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.repo.Register(reqBody.Username, reqBody.Email, reqBody.Password)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.PrepareResponse(nil, err.Error(), ""))
		return
	}

	res := s.prepareResponse(user, true)
	c.JSON(http.StatusOK, response.PrepareResponse(res, "OK", ""))
}

func (s *service) GetUser(c *gin.Context) {
	userId := c.Param("uid")
	auth := c.GetHeader("Authorization")

	user, err := s.repo.GetUser(userId, auth)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.PrepareResponse(nil, err.Error(), ""))
		return
	}
	res := s.prepareResponse(user, false)
	c.JSON(http.StatusOK, response.PrepareResponse(res, "OK", ""))
}
