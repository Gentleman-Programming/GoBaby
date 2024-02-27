package mainRoute

import (
	"GoBaby/routes"
	"GoBaby/routes/mainRoute/mainUtils"
	"GoBaby/utils"
	"fmt"
	"html/template"
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
		tmpl := template.Must(template.ParseFiles("views/main/clock.html"))
		fmt.Println(clock)
		tmpl.Execute(w, clock)
	}
}

func ClockRender() {
	mainUtils.SetDuration(duration)
	go mainUtils.StartCountdown(&clock, duration)
	routes.GetMuxInstance().HandleFunc(url, clockFragment)
}
