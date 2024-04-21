package db_config

import (
	"GoBaby/internal/models"
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
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

func InitializeDb() (*mongo.Client, *models.AppError) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, &models.AppError{
			Message: "failed to connect to MongoDB",
			Code:    500,
			Err:     err,
		}
	}

	UserCollection = InitializeUsersCollection(client)

	// check if the user with _id equal to 0 exists
	result := UserCollection.FindOne(context.TODO(), bson.M{"_id": 0})

	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			// if there is an error that means that the user has not been found so we can create a new one
			user := models.User{UserName: "test", Logs: make([]models.Log, 0), Id: 0}

			_, err := UserCollection.InsertOne(context.TODO(), user)
			if err != nil {
				return nil, &models.AppError{
					Message: "failed to create user with _id 0",
					Code:    500,
					Err:     err,
				}
			}
		} else {
			// unexpected error while querying MongoDB
			return nil, &models.AppError{
				Message: "error querying MongoDB for user with _id 0",
				Code:    500,
				Err:     err,
			}
		}
	} else {
		slog.Info("User with _id 0 exists")
	}

	slog.Info("DB STARTED")
	return client, nil
}
