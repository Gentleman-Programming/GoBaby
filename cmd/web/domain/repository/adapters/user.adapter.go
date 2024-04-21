package repository_adapters

import (
	db_config "GoBaby/cmd/web/domain/repository/config"
	"GoBaby/internal/models"
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
)

func GetUserByUUID(uuid int) (models.User, *models.AppError) {
	filter := bson.D{
		{
			Key: "$and",
			Value: bson.A{
				bson.D{
					{Key: "_id", Value: uuid},
				},
			},
		},
	}

	var result models.User
	var error *models.AppError

	err := db_config.UserCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		error = &models.AppError{
			Message: "failed to get user by uuid",
			Code:    500,
			Err:     err,
		}
	}

	return result, error
}

func SetUserByUUID(user *models.User) *models.AppError {
	_, err := db_config.UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return &models.AppError{
			Message: "failed to set user by uuid",
			Code:    500,
			Err:     err,
		}
	}

	return nil
}

func AddLogByUUID(uuid int, log models.Log) *models.AppError {
	// create a filter to find the user by uuid
	filter := bson.M{"_id": uuid}

	// create an update to push the log to the logs array
	update := bson.M{
		"$push": bson.M{
			"logs": log,
		},
	}

	slog.Info("Add Log", "log", log)

	// perform the update
	result, err := db_config.UserCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		slog.Error(err.Error())
		return &models.AppError{
			Message: "failed to add log by uuid",
			Code:    500,
			Err:     err,
		}
	}

	slog.Info("Result Log", "user", result)

	return nil // return nothing if update is successful
}
