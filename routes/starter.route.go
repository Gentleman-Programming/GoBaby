package routes

import "net/http"

var mux = http.NewServeMux()

func GetMuxInstance() *http.ServeMux {
	return mux
}
