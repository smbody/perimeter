package dao

import (
	"github.com/smbody/perimeter/model"
)

type Db interface {
	// AddUser(user *model.User) error
	GetUser(name string) (*model.User, error)
	GetUserById(id string) (*model.User, error)
	UpdateToken(token *model.AccessToken) error
	GetTokenByAccess(app string, access string) (*model.AccessToken, error)
	GetTokenByRefresh(app string, refresh string) (*model.AccessToken, error)
}

var instance Db

func Register(db Db) {
	instance = db
}

// func AddUser(user *model.User) error {
// 	return instance.AddUser(user)
// }

func GetUser(name string) (*model.User, error) {
	return instance.GetUser(name)
}

func GetUserById(id string) (*model.User, error) {
	return instance.GetUserById(id)
}

func GetTokenByAccess(app string, refresh string) (*model.AccessToken, error) {
	return instance.GetTokenByAccess(app, refresh)
}

func GetTokenByRefresh(app string, refresh string) (*model.AccessToken, error) {
	return instance.GetTokenByRefresh(app, refresh)
}

func UpdateToken(token *model.AccessToken) error {
	return instance.UpdateToken(token)
}
