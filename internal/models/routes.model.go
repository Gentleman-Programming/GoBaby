package models

type Routes struct {
	MAIN      string
	LOGS      string
	OPTIONS   string
	CLOCK     string
	LOG_TABLE string
	ERROR     string
}

var RoutesInstance = Routes{
	MAIN:      "/",
	LOGS:      "/logs",
	OPTIONS:   "/options",
	CLOCK:     "/clock",
	LOG_TABLE: "/log-table",
	ERROR:     "/error",
}
