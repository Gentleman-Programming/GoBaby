package main

import (
	"GoBaby/routes"
	"log"
	"net/http"
)

func main() {
	routes.OptionsRender()
	routes.LogRender()
	routes.MainRender()
	log.Print("Starring server on :4000")
	err := http.ListenAndServe(":4000", routes.GetMuxInstance())

	log.Fatal(err)
}
