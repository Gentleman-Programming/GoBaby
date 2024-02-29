package routes

import "net/http"

var (
	mux        = http.NewServeMux()
	fileServer = http.FileServer(http.Dir("ui/static"))
)

func GetMuxInstance() *http.ServeMux {
	return mux
}

func GetFileServerInstance() http.Handler {
	return fileServer
}
