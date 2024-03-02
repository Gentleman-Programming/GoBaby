package mainRoute

import (
	mainDomain "GoBaby/cmd/web/domain/main"
	"GoBaby/cmd/web/routes"
	"GoBaby/internal/models"
)

func MainRender() {
	routes.GetMuxInstance().HandleFunc(models.RoutesInstance.MAIN, mainDomain.MainView)
}
