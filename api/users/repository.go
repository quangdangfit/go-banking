package users

import (
	"errors"
	"github.com/google/uuid"
	"go-banking/api/accounts"
	"go-banking/database"
	"go-banking/utils/errorHandler"
	"go-banking/utils/validation"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	GenerateToken(user *User) string
	Login(username string, pass string) (*User, error)
	Register(username string, email string, pass string) (*User, error)
	GetUser(id string, jwt string) (*User, error)
}

type repo struct {
}

func NewRepository() Repository {
	return &repo{}
}

func (r *repo) GenerateToken(user *User) string {
	tokenContent := jwt.MapClaims{
		"uid":    user.UID,
		"expiry": time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	errorHandler.HandleErr(err)

	return token
}

func (r *repo) Login(username string, pass string) (*User, error) {
	// Validation before login
	valid := validation.Validate(
		[]validation.Validation{
			{Value: username, Valid: "username"},
			{Value: pass, Valid: "password"},
		})

	if valid {
		user := &User{}
		if database.DB.Where("username = ? ", username).First(&user).RecordNotFound() {
			return nil, errors.New("user not found")
		}

		// Verify password
		passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
		if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
			return nil, errors.New("wrong password")
		}

		return user, nil
	} else {
		return nil, errors.New("not valid values")
	}
}

func (r *repo) Register(username string, email string, pass string) (*User, error) {
	// Add validation to registration
	valid := validation.Validate(
		[]validation.Validation{
			{Value: username, Valid: "username"},
			{Value: email, Valid: "email"},
			{Value: pass, Valid: "password"},
		})

	if valid {
		generatedPassword := validation.HashAndSalt([]byte(pass))
		user := User{UID: uuid.New().String(), Username: username, Email: email, Password: generatedPassword}
		database.DB.Create(&user)
		accounts.NewRepository().CreateAccount(user.UID, 0)

		return &user, nil
	} else {
		return nil, errors.New("not valid values")
	}

}

func (r *repo) GetUser(uid string, jwt string) (*User, error) {
	userUUID, isValid := validation.ValidateToken(jwt)
	if isValid {
		user := &User{}
		if database.DB.Where("uid = ? ", uid).First(&user).RecordNotFound() {
			return nil, errors.New("user not found")
		}
		if user.UID == userUUID {
			return user, nil
		}
	}
	return nil, errors.New("not valid token")
}
