package repository_domain

import (
	repository_adapters "GoBaby/cmd/web/domain/repository/adapters"
	db_config "GoBaby/cmd/web/domain/repository/config"
	"GoBaby/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeBD() *mongo.Client {
	return db_config.InitializeDb()
}

func GetUserByUUID(uuid int) models.User {
	return repository_adapters.GetUserByUUID(uuid)
}

func SetUserByUUID(user *models.User) {
	repository_adapters.SetUserByUUID(user)
}

func AddLogByUUID(uuid int, log models.Log) {
	repository_adapters.AddLogByUUID(uuid, log)
}
