package routes

import (
	errorDomain "GoBaby/cmd/web/domain/error"
	"GoBaby/internal/models"
)

func ErrorRender() {
	mux.HandleFunc(models.RoutesInstance.CLEAR_ERROR, errorDomain.ClearErrorTemplate)
}
