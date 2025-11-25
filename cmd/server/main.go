package main

import (
	"log"
	"net/http"

	"github.com/doble97/scheduleapi/internal/core/ports"
	"github.com/doble97/scheduleapi/internal/core/services"
	httpHandler "github.com/doble97/scheduleapi/internal/platform/api/http/handler"
	middleware "github.com/doble97/scheduleapi/internal/platform/api/http/middlware"
	"github.com/doble97/scheduleapi/internal/platform/repository/inmemory"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	var repo ports.DashboardRepository = inmemory.NewInMemoryDashboardRepo()

	var servi = services.NewDashboardService(repo)

	dbhandler := httpHandler.NewDashboardHandler(servi)

	//Add middleware
	router.Use(middleware.LoggerMiddleware)
	// Route to create dashboard
	router.HandleFunc("/dashboard", dbhandler.CreateDashboardHandler).Methods("POST")
	router.HandleFunc("/dashboard", dbhandler.GetManyDashboardsByIDUserHandler).Methods("GET")

	serveAddr := ":8080"
	log.Printf("Starting server on %s", serveAddr)
	if err := http.ListenAndServe(serveAddr, router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}

}
