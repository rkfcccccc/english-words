package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewClient(ctx context.Context, username, password, host, port, database string) *mongo.Database {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", username, password, host, port, database)
	opts := options.Client().ApplyURI(uri)

	opts.SetAuth(options.Credential{
		Username:    username,
		Password:    password,
		PasswordSet: true,
	})

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("mongo.Connect: %v", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("client.Ping: %v", err)
	}

	return client.Database(database)
}
