package data

import (
	"errors"
	"github.com/smbody/perimeter/config"
	"github.com/smbody/perimeter/dao"
	"github.com/smbody/perimeter/model"
	"time"
)

func createToken(timeStamp time.Time, minutes time.Duration) *model.Token {
	return &model.Token{
		Token:   config.CreateToken(),
		Created: timeStamp,
		Expired: timeStamp.Add(time.Minute * minutes)}
}

// создает и сохраняет новые токены
func Login(app string, user string) *model.AccessToken {
	// сгенерировать новые токены
	ts := time.Now()

	token := &model.AccessToken{
		AppId:   app,
		UserId:  user,
		Access:  createToken(ts, config.AccessTokenLifeMinutes),
		Refresh: createToken(ts, config.RefreshTokenLifeMinutes)}

	dao.UpdateToken(token)

	return token
}

// обновляет токены по refresh
func RefreshToken(app string, refresh string) (*model.AccessToken, error) {
	// поищем токен
	token, err := dao.GetTokenByRefresh(app, refresh)
	if err != nil {
		return nil, err
	}

	// проверим его
	if !token.Refresh.Expired.After(time.Now()) {
		return nil, errors.New(config.AuthErrorBadRefreshToken)
	}

	// обновим access
	token.Access = createToken(time.Now(), config.AccessTokenLifeMinutes)
	dao.UpdateToken(token)

	return token, nil
}

// проверяет access токен
func ValidateToken(app string, access string) (*model.AccessToken, error) {
	// поищем токен
	token, err := dao.GetTokenByAccess(app, access)
	if err != nil {
		return nil, err
	}

	// проверим его
	if !token.Access.Expired.After(time.Now()) {
		return nil, errors.New(config.AuthErrorBadAccessToken)
	}

	return token, nil
}
