package routes

import (
	errorDomain "GoBaby/cmd/web/domain/error"
	"GoBaby/internal/models"
)

func ErrorRender() {
	mux.HandleFunc("GET "+models.RoutesInstance.CLEAR_ERROR, errorDomain.ClearErrorTemplate)
}
