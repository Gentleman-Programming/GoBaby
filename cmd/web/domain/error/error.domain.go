package errorDomain

import (
	repository_domain "GoBaby/cmd/web/domain/repository"
	"GoBaby/internal/models"
	"GoBaby/internal/utils"
	"GoBaby/ui"
	"fmt"
	"net/http"
)

func setErrorHeader(w http.ResponseWriter) {
	w.Header().Set("HX-Retarget", "#error")
	w.Header().Add("HX-Reswap", "outerHTML")

	w.WriteHeader(http.StatusOK)
}

func ErrorTemplate(w http.ResponseWriter, r *http.Request, errMsg *models.AppError) {
	setErrorHeader(w)

	monitorLog := models.MonitorLog{
		Err:     errMsg.Err,
		Code:    errMsg.Code,
		Message: errMsg.Message,
		Path:    r.URL.Path,
	}

	repository_domain.AddMonitorLog(monitorLog)

	file := "html/pages/error/error.tmpl.html"

	context := struct{ ErrorMessage string }{ErrorMessage: fmt.Sprint(errMsg)}
	utils.ParseTemplateFiles(w, "error", context, utils.EmptyFuncMap, ui.Content, file)
}

func ClearErrorTemplate(w http.ResponseWriter, r *http.Request) {
	setErrorHeader(w)

	file := "html/pages/error/clear-error.tmpl.html"

	utils.ParseTemplateFiles(w, "error", utils.EmptyStruct, utils.EmptyFuncMap, ui.Content, file)
}
