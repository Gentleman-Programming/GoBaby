package mainRoute

import (
	"GoBaby/routes"
	"GoBaby/utils"
	"html/template"
	"net/http"
)

var mainUrl = "/"

func mainView(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, mainUrl)

	if utils.IsValidHTTPMethod(r.Method, utils.GET.String(), w) {
		tmpl := template.Must(template.ParseFiles("views/index.html"))
		tmpl.Execute(w, nil)
	}
}

func MainRender() {
	routes.GetMuxInstance().HandleFunc("/", mainView)
}
