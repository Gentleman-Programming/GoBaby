package mainDomain

import (
	"GoBaby/internal/models"
	"GoBaby/internal/utils"
	"net/http"
)

func MainView(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.MAIN)

	files := []string{
		"html/base.html",
		"html/pages/main/main.tmpl.html",
	}

	utils.ParseTemplateFiles(w, "base", utils.EmptyStruct, utils.EmptyFuncMap, files...)
}
