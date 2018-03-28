package model

import (
	"encoding/json"
)

type User struct {
	Id       string `bson:"Id,omitempty" json:"uid"`
	name     string `bson:"name,omitempty" json:"name"`
	password string `bson:"password" json:"-"`
	passsalt string `bson:"passsalt" json:"-"`
}

func (u *User) FullName() string {
	return u.name
}

func (u *User) Password() string {
	return u.password
}

func (u *User) Salt() string {
	return u.passsalt
}

func (u *User) Json() []byte {
	if result, err := json.Marshal(u); err == nil {
		return result
	}
	return nil
}
