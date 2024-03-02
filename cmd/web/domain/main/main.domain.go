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

		context := struct{}{}
		utils.ParseTemplateFiles(w, "base", context, files...)
	}
}
