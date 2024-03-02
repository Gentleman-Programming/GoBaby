package logDomain

import (
	"GoBaby/internal/models"
	"GoBaby/internal/utils"
	"net/http"
)

func LogView(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.LOG)

	if utils.IsValidHTTPMethod(r.Method, utils.GET.String(), w) {
		w.Write([]byte("Hello, Log!"))
	}
}
