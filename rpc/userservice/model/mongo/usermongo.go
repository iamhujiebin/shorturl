package mongo

import (
	"github.com/globalsign/mgo/bson"
	"github.com/tal-tech/go-zero/core/stores/mongo"
)

type UserModel struct {
	Id       bson.ObjectId `bson:"_id"`
	UserId   int           `bson:"user_id"`
	Nickname string        `bson:"nick_name"`
	Age      int           `bson:"age"`
}

func (UserModel) DBConfig() (url, db, collection string) {
	return "localhost:27103", "FEWeb", "users"
}

func (s UserModel) FindOne(userId int) (*UserModel, error) {
	url, db, collection := s.DBConfig()
	model := mongo.MustNewModel(url, db, collection)
	session, err := model.TakeSession()
	if err != nil {
		return nil, err
	}
	var user UserModel
	err = model.GetCollection(session).Find(bson.M{"user_id": userId}).One(&user)
	if err != nil {
		return nil, err
	}
	model.PutSession(session)
	return &user, nil
}
