package models

type Routes struct {
	MAIN    string
	LOG     string
	OPTIONS string
	CLOCK   string
}

var RoutesInstance = Routes{
	MAIN:    "/",
	LOG:     "/log",
	OPTIONS: "/options",
	CLOCK:   "/clock",
}
