package db

import (
	"context"
	"fmt"
	"go-ai/pkg/util"
	"time"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DbClient *mongo.Client

func InitDB() (*mongo.Client, error) {
	dbFullHost := util.GetEnv("DB_FULL_HOST", "")
	dbPass := util.GetEnv("DB_PASSWORD", "")
	dbUser := util.GetEnv("DB_USER", "")
	host := util.GetEnv("DB_HOST", "localhost")
	dbPort := util.GetEnv("DB_PORT", "27017")

	var uri string
	if dbFullHost != "" {
		uri = fmt.Sprintf("mongodb://%v", dbFullHost)
	} else {
		if dbPass != "" {
			uri = fmt.Sprintf(`mongodb://%v:%v@%v:%v`, dbUser, dbPass, host, dbPort)
		} else {
			uri = fmt.Sprintf("mongodb://%v:%v", host, dbPort)
		}

	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	//SetComponentLevel(options.LogComponentCommand, options.LogLevelDebug)

	bsonOpts := &options.BSONOptions{
		DefaultDocumentM: true,
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMonitor(&event.CommandMonitor{
		Started: func(ctx context.Context, evt *event.CommandStartedEvent) {
		},
		Succeeded: func(ctx context.Context, evt *event.CommandSucceededEvent) {
		},
		Failed: func(ctx context.Context, evt *event.CommandFailedEvent) {
		},
	}).SetBSONOptions(bsonOpts))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	DbClient = client

	return client, nil

}
