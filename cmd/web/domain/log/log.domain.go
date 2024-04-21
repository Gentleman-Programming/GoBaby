package logDomain

import (
	errorDomain "GoBaby/cmd/web/domain/error"
	repository_domain "GoBaby/cmd/web/domain/repository"
	"GoBaby/internal/models"
	"GoBaby/internal/utils"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogViewModel struct {
	Logs []models.Log
}

func LogTable(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.LOG_TABLE)

	fmt.Println("LogTable")

	user, err := repository_domain.GetUserByUUID(0)
	if err != nil {
		errorDomain.ErrorTemplate(w, r, err)
	}

	slog.Info("LogTable", "logs", user.Logs)

	if utils.IsValidHTTPMethod(r.Method, utils.GET.String(), w) {
		utils.ParseTemplateFiles(w, "logTable", user.Logs, utils.EmptyFuncMap, "ui/html/pages/logs/logTable.tmpl.html")
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

func SaveLog(countdown int) *models.AppError {
	uuid := 0
	currentTime := time.Now()
	primitiveDateTime := primitive.NewDateTimeFromTime(currentTime)
	err := repository_domain.AddLogByUUID(uuid, models.Log{Date: primitiveDateTime, Duration: countdown})
	return err
}
