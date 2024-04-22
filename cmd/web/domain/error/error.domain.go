package errorDomain

import (
	"GoBaby/internal/models"
	"GoBaby/internal/utils"
	"fmt"
	"net/http"
)

func ErrorTemplate(w http.ResponseWriter, r *http.Request, errMsg *models.AppError) {
	w.Header().Set("HX-Retarget", "#error")
	w.Header().Add("HX-Reswap", "outerHTML")

	w.WriteHeader(http.StatusOK)
	file := "ui/html/pages/error/error.tmpl.html"

	context := struct{ ErrorMessage string }{ErrorMessage: fmt.Sprint(errMsg)}
	utils.ParseTemplateFiles(w, "error", context, utils.EmptyFuncMap, file)
}

func ClearErrorTemplate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Retarget", "#error")
	w.Header().Add("HX-Reswap", "outerHTML")

	w.WriteHeader(http.StatusOK)
	file := "ui/html/pages/error/clear-error.tmpl.html"

	utils.ParseTemplateFiles(w, "error", utils.EmptyStruct, utils.EmptyFuncMap, file)
}
