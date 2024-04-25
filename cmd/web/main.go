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
	InitRoutes()

	repository_domain.InitializeBD()

	mux := routes.GetMuxInstance()

	// serve static files
	fileServer := routes.GetFileServerInstance()
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	slog.Log(context.TODO(), slog.LevelInfo, "Starring server on :4000")
	err := http.ListenAndServe(":4000", routes.GetMuxInstance())
	log.Fatal(err)
}
