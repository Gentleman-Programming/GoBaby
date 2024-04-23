package routes

import (
	"io/fs"
	"net/http"

	"GoBaby/ui"
)

var (
	mux        = http.NewServeMux()
	static, _  = fs.Sub(ui.Content, "static")
	fileServer = http.FileServerFS(static)
)

func GetMuxInstance() *http.ServeMux {
	return mux
}

func GetFileServerInstance() http.Handler {
	return fileServer
}
