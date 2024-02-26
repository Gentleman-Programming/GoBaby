package routes

import (
	"GoBaby/utils"
	"net/http"
)

func optionsView(w http.ResponseWriter, r *http.Request) {
	if utils.IsValidHTTPMethod(r.Method, utils.GET.String(), w) {
		w.Write([]byte("Hello, Options!"))
	}
}

func OptionsRender() {
	mux.HandleFunc("/options", optionsView)
}
