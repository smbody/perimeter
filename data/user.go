package data

import (
	"errors"

	"github.com/smbody/perimeter/config"
	"github.com/smbody/perimeter/dao"
	"github.com/smbody/perimeter/model"
)

// найдем пользователя по логин/пароль
func GetUser(app string, name string, pass string) (*model.User, error) {
	user, err := dao.GetUser(name)
	if err != nil {
		return nil, err
	}

	password := pass + user.Salt() // подсолим
	// проверим пароль
	if !config.PasswordValidate(user.Password, password) {
		return nil, errors.New(config.AuthErrorBadUserPassword)
	}

	return user, nil
}

// найдем пользователя по id
func GetUserById(app string, id string) (*model.User, error) {
	return dao.GetUserById(id)
}

// добавление нового пользователя
func CreateUser(app string, name string, pass string) (*model.User, error) {
	// пароль должен быть корректным
	if len(pass)+config.SaltSize > config.PasswordSize || len(pass) == 0 {
		return nil, errors.New(config.AuthErrorBadUserPassword)
	}

	salt := config.CreateSalt()
	password, err := config.GetPasswordHash(pass + salt) // подсолим
	if err != nil {
		return nil, errors.New(config.AuthErrorBadUserPassword)
	}

	user := model.CreateUser(name, password, salt)

	user, err = dao.AddUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
