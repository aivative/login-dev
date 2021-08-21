package utils

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

// NewMongoConnection connect to mongodb server, returning a reusable mongo connection.
func NewMongoConnection(ctx context.Context, uri string) (client *mongo.Client, err error) {
	clientOpt := options.
		Client().
		ApplyURI(uri).
		SetWriteConcern(writeconcern.New(writeconcern.WMajority())).
		// SetReadConcern(readconcern.Snapshot()).
		SetReadConcern(readconcern.Majority()).
		SetReadPreference(readpref.Primary())

	if client, err = mongo.NewClient(clientOpt); err != nil {
		return
	}

	if err = client.Connect(ctx); err != nil {
		return
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return
	}

	return client, nil
}

// NotFound is shortcut for Error no documents in mongodb
func NotFound(err error) bool {
	return err == mongo.ErrNoDocuments
}
