package mainDomain

import (
	errorDomain "GoBaby/cmd/web/domain/error"
	logDomain "GoBaby/cmd/web/domain/log"
	"GoBaby/internal/models"
	"GoBaby/internal/utils"
	"net/http"
)

var duration = 14400

var ClockInstance = utils.NewClock()

func GetClock() *utils.Clock {
	return ClockInstance
}

func GetDuration() int {
	return duration
}

func ClockFragment(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.CLOCK)

	if utils.IsValidHTTPMethod(r.Method, utils.GET.String(), w) {
		utils.ParseTemplateFiles(w, "clock", ClockInstance, utils.EmptyFuncMap, "ui/html/pages/main/clock.tmpl.html")
	}
}

func RestartCycle(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.MAIN)

	if utils.IsValidHTTPMethod(r.Method, utils.POST.String(), w) {
		select {
		case <-ClockInstance.Stop: // If the channel is already closed, do nothing
		default:
			utils.StopCountdown(ClockInstance)
		}

		err := logDomain.SaveLog(utils.FormatCountdownToTimestamp(ClockInstance.CountDown))
		if err != nil {
			errorDomain.ErrorTemplate(w, r, err)
		}

		ClockInstance.CountDown = "04:00:00"

		utils.SetDuration(duration)
		go utils.StartCountdown(ClockInstance, duration)
	}
}
