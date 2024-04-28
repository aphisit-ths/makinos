package config

import (
	"context"
	"fmt"
	"makino/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToCustomerDB(setting models.AppSettings) (*mongo.Database, context.Context, *mongo.Client, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 2500*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(setting.ConnectionStrings.CustomerDocumentDb))
	if err != nil {
		cancel()
		panic("failed to connect MongoDb ")
	}
	fmt.Println("### Connected to MongoDb Successfully")
	mongoDatabase := client.Database(setting.ConnectionStrings.CustomerDatabaseName)

	return mongoDatabase, ctx, client, cancel
}
