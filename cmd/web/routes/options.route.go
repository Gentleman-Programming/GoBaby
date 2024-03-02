package routes

import (
	optionsDomain "GoBaby/cmd/web/domain/options"
	"GoBaby/internal/models"
)

func OptionsRender() {
	mux.HandleFunc(models.RoutesInstance.OPTIONS, optionsDomain.OptionsView)
}
