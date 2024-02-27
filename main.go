package main

import (
	"GoBaby/routes"
	"GoBaby/routes/mainRoute"
	"log"
	"net/http"
)

func main() {
	routes.OptionsRender()
	routes.LogRender()
	mainRoute.MainRender()
	mainRoute.ClockRender()
	log.Print("Starring server on :4000")
	err := http.ListenAndServe(":4000", routes.GetMuxInstance())

	log.Fatal(err)
}
