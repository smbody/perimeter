package dao

import (
	"fmt"
	"github.com/smbody/perimeter/config"
	"github.com/smbody/perimeter/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"sync"
)

type mongoDb struct {
	Session *mgo.Session
	DB      *mgo.Database
}

const (
	mongoHost        = "perimeter-data:28018"
	mongoDB          = "perimeter"
	collectionUsers  = "users"
	collectionTokens = "tokens"
)

var mongo *mongoDb
var once sync.Once

func (c *mongoDb) init() error {
	fmt.Println("Dialing to mongo", mongoHost)
	session, err := mgo.Dial(mongoHost)
	if err != nil {
		return err
	}
	fmt.Printf("Mongo status ok!")

	c.Session = session
	c.Session.SetMode(mgo.Monotonic, true)
	c.DB = c.Session.DB(mongoDB)

	return nil
}

func Mongo() *mongoDb {
	once.Do(func() {
		mongo = &mongoDb{}
		if err := mongo.init(); err != nil {
			panic(err)
		}
	})

	if err := mongo.Session.Ping(); err != nil {
		mongo.Session.Refresh()
	}
	return mongo
}

func getUsers() *mgo.Collection {
	return Mongo().DB.C(collectionUsers)
}

func getTokens() *mgo.Collection {
	return Mongo().DB.C(collectionTokens)
}

func (db *mongoDb) GetUser(name string) (*model.User, error) {
	var user model.User
	err := getUsers().Find(bson.M{"name": name}).One(&user)
	return &user, err
}

func (db *mongoDb) GetUserById(id string) (*model.User, error) {
	var user model.User
	err := getUsers().Find(bson.M{"_id": id}).One(&user)
	return &user, err
}

func (db *mongoDb) GetTokenByAccess(app string, access string) (*model.AccessToken, error) {
	var token model.AccessToken
	err := getTokens().Find(bson.M{"aid": app, "access.token": access}).One(&token)

	return &token, err
}

func (db *mongoDb) GetTokenByRefresh(app string, refresh string) (*model.AccessToken, error) {
	var token model.AccessToken
	err := getTokens().Find(bson.M{"aid": app, "refresh.token": refresh}).One(&token)

	return &token, err
}

func (db *mongoDb) UpdateToken(token *model.AccessToken) error {
	_, err := getTokens().Upsert(
		bson.M{"aid": token.AppId, "uid": token.UserId},
		bson.M{"aid": token.AppId, "uid": token.UserId, "access": token.Access, "refresh": token.Refresh})

	return err
}

func (db *mongoDb) AddUser(user *model.User) (*model.User, error) {
	// id сами делаем, потому что используем string
	user.Id = config.GetHash(user.FullName() + user.Salt())

	err := getUsers().Insert(user)
	if err != nil {
		return nil, err
	}

	return db.GetUser(user.Name)
}
