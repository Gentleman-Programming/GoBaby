package routes

import (
	optionsDomain "GoBaby/cmd/web/domain/options"
	"GoBaby/internal/models"
)

func OptionsRender() {
	mux.HandleFunc("GET "+models.RoutesInstance.OPTIONS, optionsDomain.OptionsView)
}
