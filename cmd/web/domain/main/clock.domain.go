package mainDomain

import (
	"GoBaby/internal/models"
	"GoBaby/internal/utils"
	"fmt"
	"net/http"
)

var duration = 14400

var clock = utils.Clock{
	Stop:      make(chan struct{}, 1), // Channel for stop signal (buffered to prevent blocking)
	CountDown: "04:00:00",
}

func GetClock() *utils.Clock {
	return &clock
}

func GetDuration() int {
	return duration
}

func ClockFragment(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.CLOCK)

	if utils.IsValidHTTPMethod(r.Method, utils.GET.String(), w) {
		files := []string{
			"ui/html/pages/main/clock.tmpl.html",
		}

		utils.ParseTemplateFiles(w, "clock", clock, files...)
	}
}

func RestartCycle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Restarting cycle")
	utils.CheckIfPath(w, r, models.RoutesInstance.MAIN)

	if utils.IsValidHTTPMethod(r.Method, utils.POST.String(), w) {
		select {
		case <-clock.Stop: // If the channel is already closed, do nothing
		default:
			utils.StopCountdown(&clock)
		}

		clock.CountDown = "04:00:00"

		utils.SetDuration(duration)
		go utils.StartCountdown(&clock, duration)
	}
}
