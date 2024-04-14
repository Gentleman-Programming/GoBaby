package db_config

import (
	"GoBaby/internal/models"
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database = "GoBaby"

var collections = map[string]string{
	"users": "users",
	"logs":  "logs",
}

var UserCollection *mongo.Collection

func InitializeUsersCollection(client *mongo.Client) *mongo.Collection {
	return client.Database(database).Collection(collections["users"])
}

func InitializeLogsCollection(client *mongo.Client) *mongo.Collection {
	return client.Database(database).Collection(collections["logs"])
}

func InitializeDb() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	UserCollection = InitializeUsersCollection(client)

	user := models.User{Id: 0, UserName: "test", Logs: make([]models.Log, 0)}

	UserCollection.InsertOne(context.TODO(), user)

	slog.Info("DB STARTED")
	return client
}
