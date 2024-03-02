package routes

import (
	logDomain "GoBaby/cmd/web/domain/log"
	"GoBaby/internal/models"
)

func LogRender() {
	mux.HandleFunc(models.RoutesInstance.LOG, logDomain.LogView)
}
