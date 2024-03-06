package routes

import (
	logDomain "GoBaby/cmd/web/domain/log"
	"GoBaby/internal/models"
)

func LogRender() {
	mux.HandleFunc(models.RoutesInstance.LOG_TABLE, logDomain.LogTable)
	mux.HandleFunc(models.RoutesInstance.LOGS, logDomain.LogView)
}
