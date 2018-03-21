package data

import (
// "github.com/smbody/perimeter/config"
)

// var tokens

func CreateAccessToken(app string, user *User) string {
	return "frctcnjrty"
}

func ValidateToken(app string, token string) (*User, string) {
	user, err := GetUser(app, "User for token: "+token, "")
	if err != nil {
		return nil, err.Error()
	}
	user.Id = token

	return user, ""
}
