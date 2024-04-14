package logDomain

import (
	repository_domain "GoBaby/cmd/web/domain/repository"
	"GoBaby/internal/models"
	"GoBaby/internal/utils"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type LogViewModel struct {
	Logs []models.Log
}

func LogTable(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.LOG_TABLE)

	fmt.Println("LogTable")

	logs := repository_domain.GetUserByUUID(0).Logs

	slog.Info("LogTable", "logs", logs)

	if utils.IsValidHTTPMethod(r.Method, utils.GET.String(), w) {
		utils.ParseTemplateFiles(w, "logTable", logs, utils.EmptyFuncMap, "ui/html/pages/logs/logTable.tmpl.html")
	}
}

func LogView(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.LOGS)

	if utils.IsValidHTTPMethod(r.Method, utils.GET.String(), w) {
		files := []string{
			"ui/html/base.html",
			"ui/html/pages/logs/logs.tmpl.html",
		}

		utils.ParseTemplateFiles(w, "base", utils.EmptyStruct, utils.EmptyFuncMap, files...)
	}
}

func SaveLog(countdown int) {
	uuid := 0
	repository_domain.AddLogByUUID(uuid, models.Log{Date: time.Now(), Duration: countdown})
}
