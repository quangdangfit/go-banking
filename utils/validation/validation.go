package validation

import (
	"github.com/dgrijalva/jwt-go"
	"go-banking/utils/errorHandler"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strings"
)

type Validation struct {
	Value string
	Valid string
}

func Validate(values []Validation) bool {
	username := regexp.MustCompile("[A-Za-z0-9]")
	email := regexp.MustCompile("^[A-Za-z0-9]+[@]+[A-Za-z0-9]+[.]+[A-Za-z]+$")

	for i := 0; i < len(values); i++ {
		switch values[i].Valid {
		case "username":
			if !username.MatchString(values[i].Value) {
				return false
			}
		case "email":
			if !email.MatchString(values[i].Value) {
				return false
			}
		case "password":
			if len(values[i].Value) < 5 {
				return false
			}
		}
	}
	return true
}

func ValidateToken(jwtToken string) (string, bool) {
	if jwtToken == "" {
		return "", false
	}
	cleanJWT := strings.Replace(jwtToken, "Bearer ", "", -1)
	tokenData := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(cleanJWT, tokenData, func(token *jwt.Token) (interface{}, error) {
		return []byte("TokenPassword"), nil
	})
	errorHandler.HandleErr(err)

	if token.Valid {
		return tokenData["uid"].(string), true
	} else {
		return "", false
	}
}

func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	errorHandler.HandleErr(err)

	return string(hashed)
}
