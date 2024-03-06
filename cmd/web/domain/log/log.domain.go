package logDomain

import (
	"GoBaby/internal/models"
	"GoBaby/internal/utils"
	"fmt"
	"net/http"
	"time"
)

type LogViewModel struct {
	Logs []models.Log
}

var logs = []models.Log{
	{Date: time.Now(), Duration: 10}, {Date: time.Now().Add(time.Hour), Duration: 15},
}

func LogTable(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.LOG_TABLE)

	fmt.Println("LogTable")

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
	logs = append(logs, models.Log{Date: time.Now(), Duration: countdown})
}
