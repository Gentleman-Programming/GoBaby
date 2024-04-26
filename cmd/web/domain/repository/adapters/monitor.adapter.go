package repository_adapters

import (
	db_config "GoBaby/cmd/web/domain/repository/config"
	"GoBaby/internal/models"
	"context"
	"fmt"
)

func AddMonitorLog(monitorLog models.MonitorLog) *models.AppError {
	_, err := db_config.MonitorCollection.InsertOne(context.TODO(), monitorLog)
	if err != nil {
		return &models.AppError{
			Message: "Error updating Monitor",
			Err:     err,
			Code:    500,
		}
	}

	fmt.Println("Monitor Log added", monitorLog)
	return nil
}
