package main

import (
	"log"
	"net/http"

	"github.com/doble97/scheduleapi/config"
	"github.com/doble97/scheduleapi/internal/app"
	"github.com/doble97/scheduleapi/internal/platform/api/http/router"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}
	appContext := app.NewAppContext(*config.Config)
	// appContext.
	router := router.NewAPIRouter(appContext)

	serveAddr := ":8080"
	log.Printf("Starting server on %s", serveAddr)
	if err := http.ListenAndServe(serveAddr, router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}

}
