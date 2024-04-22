package mainDomain

import (
	errorDomain "GoBaby/cmd/web/domain/error"
	logDomain "GoBaby/cmd/web/domain/log"
	"GoBaby/internal/models"
	"GoBaby/internal/utils"
	"net/http"
)

var duration = 14400

var clockInstance = utils.NewClock()

func GetClock() *utils.Clock {
	return clockInstance
}

func GetDuration() int {
	return duration
}

func ClockFragment(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.CLOCK)

	utils.ParseTemplateFiles(w, "clock", clockInstance, utils.EmptyFuncMap, "ui/html/pages/main/clock.tmpl.html")
}

func RestartCycle(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.RESTART_CYCLE)

	select {
	case <-clockInstance.Stop: // If the channel is already closed, do nothing
	default:
		err := logDomain.SaveLog(utils.FormatCountdownToTimestamp(clockInstance.CountDown))
		if err != nil {
			errorDomain.ErrorTemplate(w, r, err)
			return
		}

		utils.StopCountdown(clockInstance)
		clockInstance.CountDown = "04:00:00"
		utils.SetDuration(duration)
		go utils.StartCountdown(clockInstance, duration)

		w.Write([]byte("Cycle restarted"))
	}
}
