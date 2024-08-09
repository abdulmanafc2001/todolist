package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(pass), nil
}

func VerifyPassword(pass string, userPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPass), []byte(pass))
	return err == nil
}

var secret_key = os.Getenv("SECRET_KEY")

func CreateToken(userName, userType string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
		"sub":      userName,
		"usertype": userType,
	})
	return token.SignedString([]byte(secret_key))
}
