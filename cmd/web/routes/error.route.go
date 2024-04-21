package routes

import (
	errorDomain "GoBaby/cmd/web/domain/error"
	"GoBaby/internal/models"
	"net/http"
)

func ErrorRender(errorMessage *models.AppError) {
	mux.HandleFunc(models.RoutesInstance.ERROR, func(w http.ResponseWriter, r *http.Request) {
		errorDomain.ErrorTemplate(w, r, errorMessage)
	})
}
