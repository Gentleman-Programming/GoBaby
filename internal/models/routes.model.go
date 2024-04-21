package models

type Routes struct {
	MAIN          string
	LOGS          string
	OPTIONS       string
	CLOCK         string
	RESTART_CYCLE string
	LOG_TABLE     string
	ERROR         string
}

var RoutesInstance = Routes{
	MAIN:          "/",
	LOGS:          "/logs",
	OPTIONS:       "/options",
	CLOCK:         "/clock",
	RESTART_CYCLE: "/clock/restart-cycle",
	LOG_TABLE:     "/log-table",
	ERROR:         "/error",
}
