package mainRoute

import (
	"GoBaby/cmd/web/routes"
	"GoBaby/cmd/web/routes/mainRoute/mainUtils"
	"GoBaby/internal/utils"
	"fmt"
	"net/http"
)

var url = "/clock"

var duration = 14400

var clock = mainUtils.Clock{
	Stop:      make(chan struct{}, 1), // Channel for stop signal (buffered to prevent blocking)
	CountDown: "04:00:00",
}

func clockFragment(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, url)

	if utils.IsValidHTTPMethod(r.Method, utils.GET.String(), w) {
		files := []string{
			"ui/html/pages/main/clock.tmpl.html",
		}

		utils.ParseTemplateFiles(w, "clock", clock, files...)
	}
}

func restartCycle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Restarting cycle")
	utils.CheckIfPath(w, r, mainUrl)

	if utils.IsValidHTTPMethod(r.Method, utils.POST.String(), w) {
		select {
		case <-clock.Stop: // Check if already stopped
		default:
			mainUtils.StopCountdown(&clock)
		}

		clock.CountDown = "04:00:00"
		mainUtils.SetDuration(duration)
		go mainUtils.StartCountdown(&clock, duration)
	}
}

func ClockRender() {
	mainUtils.SetDuration(duration)
	go mainUtils.StartCountdown(&clock, duration)
	routes.GetMuxInstance().HandleFunc(url, clockFragment)
	routes.GetMuxInstance().HandleFunc("/clock/restart-cycle", restartCycle)
}
