package model

import (
	"time"
)

type Token struct {
	Token   string    `bson:"token,omitempty" json:"token"`
	Created time.Time `bson:"created" json:"created"`
	Expired time.Time `bson:"expired" json:"expired"`
}

type AccessToken struct {
	AppId   string `bson:"aid,omitempty" json:"-"`
	UserId  string `bson:"uid,omitempty" json:"-"`
	Access  *Token `bson:"access,omitempty" json:"access"`
	Refresh *Token `bson:"refresh" json:"refresh"`
}
