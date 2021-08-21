package repository

import (
	"context"

	"github.com/aivative/login-dev/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewAuthRepo(client *mongo.Client) *AuthMongoRepo {
	mongoConfig := config.MongoConf["mongo-user-auth"]

	collection := client.Database(mongoConfig.DBName).Collection(mongoConfig.CollName)
	return &AuthMongoRepo{
		coll: collection,
	}
}

type AuthMongoRepo struct {
	coll *mongo.Collection
}

type TAuthSecret struct {
	Name   string `json:"name" bson:"name"`
	Secret string `json:"secret" bson:"secret"`
}

func (amr *AuthMongoRepo) GetUserPasswordSecret(ctx context.Context) (result string, err error) {
	var authSecret TAuthSecret
	err = amr.coll.FindOne(ctx, bson.M{"name": "USER_PASSWORD_SECRET"}).Decode(&authSecret)
	return authSecret.Secret, err
}

func (amr *AuthMongoRepo) GetUserAccountSecret(ctx context.Context) (result string, err error) {
	var authSecret TAuthSecret
	err = amr.coll.FindOne(ctx, bson.M{"name": "USER_ACCOUNT_SECRET"}).Decode(&authSecret)
	return authSecret.Secret, err
}
