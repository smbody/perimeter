package data

import (
	"errors"
	"time"

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

	// проверим пароль
	password := config.GetPasswordHash(pass + user.Salt())
	if user.Password() != password {
		return nil, errors.New(config.AuthErrorBadUserPassword)
	}

	return user, nil
}

// найдем пользователя по id
func GetUserById(app string, id string) (*model.User, error) {
	return dao.GetUserById(id)
}

// создает и сохраняет новые токены
func Login(app string, user string) *model.AccessToken {
	// сгенерировать новые токены
	ts := time.Now()

	access := &model.Token{
		Token:   config.CreateToken(),
		Created: ts,
		Expired: ts.Add(time.Minute * config.AccessTokenLifeMinutes)}

	refresh := &model.Token{
		Token:   config.CreateToken(),
		Created: ts,
		Expired: ts.Add(time.Minute * config.AccessTokenLifeMinutes)}

	token := &model.AccessToken{
		AppId:   app,
		UserId:  user,
		Access:  *access,
		Refresh: *refresh}

	dao.UpdateToken(token)

	return token
}
