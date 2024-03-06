package mainDomain

import (
	"GoBaby/internal/models"
	"GoBaby/internal/utils"
	"net/http"
)

func MainView(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.MAIN)

	if utils.IsValidHTTPMethod(r.Method, utils.GET.String(), w) {
		files := []string{
			"ui/html/base.html",
			"ui/html/pages/main/main.tmpl.html",
		}

		utils.ParseTemplateFiles(w, "base", utils.EmptyStruct, utils.EmptyFuncMap, files...)
	}
}
