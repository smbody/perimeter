package data

import (
	"encoding/json"
)

type User struct {
	Id   string
	name string
}

func (u *User) FullName() string {
	return u.name
}

func (u *User) Json() []byte {
	if result, err := json.Marshal(u); err == nil {
		return result
	}
	return nil
}

func (u *User) AccessToken(app string) string {
	return app
}

func (u *User) RefreshToken(app string) string {
	return app
}

func (u *User) Login(app string) {

}

func RefreshTokens(app string, token string) (*User, error) {
	return GetUser(app, token, "")
}

func GetUser(app string, name string, pass string) (*User, error) {
	user := User{Id: app, name: name}

	return &user, nil
}
