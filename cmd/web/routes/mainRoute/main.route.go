package mainRoute

import (
	"GoBaby/cmd/web/routes"
	"GoBaby/internal/utils"
	"net/http"
)

var mainUrl = "/"

func mainView(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, mainUrl)

	if utils.IsValidHTTPMethod(r.Method, utils.GET.String(), w) {
		files := []string{
			"ui/html/base.html",
			"ui/html/pages/main/main.tmpl.html",
		}

		context := struct{}{}
		utils.ParseTemplateFiles(w, "base", context, files...)
	}
}

func MainRender() {
	routes.GetMuxInstance().HandleFunc("/", mainView)
}
