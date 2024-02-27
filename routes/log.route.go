package routes

import (
	"GoBaby/utils"
	"net/http"
)

var logUrl = "/log"

func logView(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, logUrl)

	if utils.IsValidHTTPMethod(r.Method, utils.GET.String(), w) {
		w.Write([]byte("Hello, Log!"))
	}
}

func LogRender() {
	mux.HandleFunc(logUrl, logView)
}
