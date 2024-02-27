package routes

import (
	"GoBaby/utils"
	"net/http"
)

var optionsUrl = "/options"

func optionsView(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, optionsUrl)

	if utils.IsValidHTTPMethod(r.Method, utils.GET.String(), w) {
		w.Write([]byte("Hello, Options!"))
	}
}

func OptionsRender() {
	mux.HandleFunc(optionsUrl, optionsView)
}
