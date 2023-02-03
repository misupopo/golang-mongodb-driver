package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/xerrors"
)

type SetupData interface {
	MongodbDsn() string
	MongodbDsnDBName() string
}

type MongoDB struct {
	dsn string
	*mongo.Client
	*mongo.Database
}

func NewMongoDB(d SetupData) (*MongoDB, error) {
	dsn := d.MongodbDsn()
	databaseName := d.MongodbDsnDBName()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, xerrors.Errorf("mongodb connect open error: %w", err)
	}

	database := client.Database(databaseName)

	return &MongoDB{
		dsn:      dsn,
		Client:   client,
		Database: database,
	}, nil
}

func (d *MongoDB) Close() {
	err := d.Client.Disconnect(context.TODO())
	if err != nil {
		fmt.Printf("mongodb close error: %+v", err)
	}
}
