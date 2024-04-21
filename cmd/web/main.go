package main

import (
	repository_domain "GoBaby/cmd/web/domain/repository"
	"GoBaby/cmd/web/routes"
	"GoBaby/cmd/web/routes/mainRoute"
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	routes.OptionsRender()
	routes.LogRender()
	mainRoute.MainRender()
	mainRoute.ClockRender()

	_, initDbError := repository_domain.InitializeBD()
	if initDbError != nil {
		slog.Error(fmt.Sprint(initDbError))
	}

	mux := routes.GetMuxInstance()
	fileServer := routes.GetFileServerInstance()

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	slog.Log(context.TODO(), slog.LevelInfo, "Starring server on :4000")

	err := http.ListenAndServe(":4000", routes.GetMuxInstance())
	log.Fatal(err)
}
