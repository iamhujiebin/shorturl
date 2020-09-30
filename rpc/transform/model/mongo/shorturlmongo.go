package mongo

import (
	"github.com/globalsign/mgo/bson"
	"github.com/tal-tech/go-zero/core/stores/mongo"
)

type ShortUrlModel struct {
	Id      bson.ObjectId `bson:"_id"`
	Shorten string        `bson:"shorten"`
	Url     string        `bson:"url"`
	Count   int           `bson:"count"`
}

func (ShortUrlModel) DBConfig() (url, db, collection string) {
	return "localhost:27103", "ShortUrl", "short_urls"
}

func (s ShortUrlModel) FindOne(shorten string) (*ShortUrlModel, error) {
	url, db, collection := s.DBConfig()
	model := mongo.MustNewModel(url, db, collection)
	session, err := model.TakeSession()
	if err != nil {
		return nil, err
	}
	var shortUrl ShortUrlModel
	err = model.GetCollection(session).Find(bson.M{"shorten": shorten}).One(&shortUrl)
	if err != nil {
		return nil, err
	}
	model.PutSession(session)
	return &shortUrl, nil
}
