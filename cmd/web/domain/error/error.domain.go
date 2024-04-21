package errorDomain

import (
	"GoBaby/internal/models"
	"GoBaby/internal/utils"
	"fmt"
	"net/http"
)

func ErrorTemplate(w http.ResponseWriter, r *http.Request, errMsg *models.AppError) {
	fmt.Println("ErrorTemplate")

	file := "ui/html/pages/error/error.tmpl.html"

	context := struct{ ErrorMessage string }{ErrorMessage: fmt.Sprint(errMsg)}
	utils.ParseTemplateFiles(w, "error", context, utils.EmptyFuncMap, file)
}
