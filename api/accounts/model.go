package accounts

import (
	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	AccNumber string
	UUID      string
	Type      string
	Balance   uint
	UserUUID  string
}

type AccountResponse struct {
	UUID     string
	UserUUID string
	Name     string
	Balance  int
}
