package routes

import (
	"GoBaby/utils"
	"net/http"
)

func logView(w http.ResponseWriter, r *http.Request) {
	if utils.IsValidHTTPMethod(r.Method, utils.GET.String(), w) {
		w.Write([]byte("Hello, Log!"))
	}
}

func LogRender() {
	mux.HandleFunc("/log", logView)
}
