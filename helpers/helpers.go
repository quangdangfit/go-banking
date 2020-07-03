package helpers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"
	"gitlab.com/quangdangfit/gocommon/utils/logger"
	"go-banking/interfaces"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func HandleErr(err error) {
	if err != nil {
		logger.Error(err.Error())
	}
}

func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	HandleErr(err)

	return string(hashed)
}

func Validation(values []interfaces.Validation) bool {
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

func PanicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			error := recover()
			if error != nil {
				log.Println(error)

				resp := interfaces.ErrResponse{Message: "Internal server error"}
				json.NewEncoder(w).Encode(resp)
			}
		}()
		next.ServeHTTP(w, r)
	})
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
	HandleErr(err)

	if token.Valid {
		return tokenData["uid"].(string), true
	} else {
		return "", false
	}
}

func ReadBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	HandleErr(err)

	return body
}

func PrepareResponse(data interface{}, message string, code string) map[string]interface{} {
	result := map[string]interface{}{
		"data":    data,
		"message": message,
	}

	return result
}
