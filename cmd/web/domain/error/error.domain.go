package errorDomain

import (
	"GoBaby/internal/models"
	"GoBaby/internal/utils"
	"fmt"
	"net/http"
)

func ErrorTemplate(w http.ResponseWriter, r *http.Request, errMsg *models.AppError) {
	utils.CheckIfPath(w, r, models.RoutesInstance.ERROR)

	fmt.Println("ErrorTemplate")

	files := []string{
		"ui/html/base.html",
		"ui/html/pages/error/error.tmpl.html",
	}
	if utils.IsValidHTTPMethod(r.Method, utils.GET.String(), w) {
		context := struct{ ErrorMessage *models.AppError }{ErrorMessage: errMsg}

		utils.ParseTemplateFiles(w, "error", context, utils.EmptyFuncMap, files...)
	}
}
