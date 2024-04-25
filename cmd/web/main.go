package main

import (
	repository_domain "GoBaby/cmd/web/domain/repository"
	"GoBaby/cmd/web/routes"
	"GoBaby/cmd/web/routes/mainRoute"
	"context"
	"log"
	"log/slog"
	"net/http"
)

func InitRoutes() {
	routes.OptionsRender()
	routes.LogRender()
	routes.ErrorRender()
	mainRoute.MainRender()
	mainRoute.ClockRender()
}

func main() {
	// init routes and database
	InitRoutes()
	repository_domain.InitializeBD()

	// serve static files
	mux := routes.GetMuxInstance()
	fileServer := routes.GetFileServerInstance()
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// start server
	slog.Log(context.TODO(), slog.LevelInfo, "Starring server on :4000")
	err := http.ListenAndServe(":4000", routes.GetMuxInstance())
	if err != nil {
		log.Fatal(err)
	}
}
