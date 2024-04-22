package routes

import (
	logDomain "GoBaby/cmd/web/domain/log"
	"GoBaby/internal/models"
)

func LogRender() {
	mux.HandleFunc("GET "+models.RoutesInstance.LOG_TABLE, logDomain.LogTable)
	mux.HandleFunc("GET "+models.RoutesInstance.LOGS, logDomain.LogView)
}
