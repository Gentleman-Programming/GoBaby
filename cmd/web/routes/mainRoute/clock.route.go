package mainRoute

import (
	"GoBaby/cmd/web/routes"
	"GoBaby/cmd/web/routes/mainRoute/mainUtils"
	"GoBaby/internal/utils"
	"net/http"
)

var url = "/clock"

var duration = 14400

var clock = mainUtils.Clock{
	CountDown: "04:00:00",
}

func clockFragment(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, url)

	if utils.IsValidHTTPMethod(r.Method, utils.GET.String(), w) {
		files := []string{
			"ui/html/pages/main/clock.tmpl.html",
		}

		utils.ParseTemplateFiles(w, files...)
	}
}

func ClockRender() {
	mainUtils.SetDuration(duration)
	go mainUtils.StartCountdown(&clock, duration)
	routes.GetMuxInstance().HandleFunc(url, clockFragment)
}
