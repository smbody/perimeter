package data

import (
	"github.com/dgrijalva/jwt-go"
)

// var tokens

func SaveToken(app string, user *User, token *jwt.Token) {

}

func ValidateToken(app string, token *jwt.Token) bool {
	return true
}

func CreateAccessToken(app string, user *User) string {
	return "frctcnjrty"
}
