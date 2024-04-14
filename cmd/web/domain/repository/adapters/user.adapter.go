package repository_adapters

import (
	db_config "GoBaby/cmd/web/domain/repository/config"
	"GoBaby/internal/models"
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
)

func GetUserByUUID(uuid int) models.User {
	filter := bson.D{
		{
			Key: "$and",
			Value: bson.A{
				bson.D{
					{Key: "Id", Value: uuid},
				},
			},
		},
	}

	var result models.User

	err := db_config.UserCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		slog.Error(err.Error())
	}

	return result
}

func SetUserByUUID(user *models.User) {
	_, err := db_config.UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		slog.Error(err.Error())
	}
}

func AddLogByUUID(uuid int, log models.Log) {
	filter := bson.M{"Id": uuid}

	update := bson.M{
		"$push": bson.M{
			"logs": log,
		},
	}

	slog.Info("Add Log", "log", log)

	all, _ := db_config.UserCollection.Find(context.TODO(), bson.D{})
	slog.Info("Result Log All", "All", all)

	result, err := db_config.UserCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info("Result Log", "user", result)
}
