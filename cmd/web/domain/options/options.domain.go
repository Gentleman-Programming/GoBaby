package optionsDomain

import (
	"GoBaby/internal/models"
	"GoBaby/internal/utils"
	"net/http"
)

func OptionsView(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.OPTIONS)

	if utils.IsValidHTTPMethod(r.Method, utils.GET.String(), w) {
		w.Write([]byte("Hello, Options!"))
	}
}
