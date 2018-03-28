package data

import (
	"errors"
	"time"

	"github.com/smbody/perimeter/config"
	"github.com/smbody/perimeter/dao"
	"github.com/smbody/perimeter/model"
)

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
	ts := time.Now()
	access := &model.Token{
		Token:   config.CreateToken(),
		Created: ts,
		Expired: ts.Add(time.Minute * config.AccessTokenLifeMinutes)}
	token.Access = *access

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
