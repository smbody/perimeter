package model

import (
	"encoding/json"
)

type User struct {
	Id       string `bson:"_id" json:"uid"`
	Name     string `bson:"name" json:"name"`
	Password string `bson:"password" json:"-"`
	Passsalt string `bson:"passsalt" json:"-"`
}

func (u *User) FullName() string {
	return u.Name
}

func (u *User) Salt() string {
	return u.Passsalt
}

func (u *User) Json() []byte {
	if result, err := json.Marshal(u); err == nil {
		return result
	}
	return nil
}

func CreateUser(name string, pass string, salt string) *User {
	return &User{
		Name:     name,
		Password: pass,
		Passsalt: salt}
}
