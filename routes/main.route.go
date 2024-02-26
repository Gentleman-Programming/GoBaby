package routes

import (
	"GoBaby/utils"
	"net/http"
)

func mainView(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if utils.IsValidHTTPMethod(r.Method, utils.GET.String(), w) {
		w.Write([]byte("Hello, Main!"))
	}
}

func MainRender() {
	mux.HandleFunc("/", mainView)
}
