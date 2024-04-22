package mainRoute

import (
	mainDomain "GoBaby/cmd/web/domain/main"
	"GoBaby/cmd/web/routes"
	"GoBaby/internal/models"
	"GoBaby/internal/utils"
)

func ClockRender() {
	duration := mainDomain.GetDuration()
	clock := mainDomain.GetClock()
	clockFragment := mainDomain.ClockFragment
	restartCycle := mainDomain.RestartCycle

	utils.SetDuration(duration)
	go utils.StartCountdown(clock, duration)

	// Routes
	routes.GetMuxInstance().HandleFunc(models.RoutesInstance.CLOCK, clockFragment)
	routes.GetMuxInstance().HandleFunc("POST "+models.RoutesInstance.RESTART_CYCLE, restartCycle)
}
