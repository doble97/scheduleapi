package router

import (
	"github.com/doble97/scheduleapi/internal/app"
	middleware "github.com/doble97/scheduleapi/internal/platform/api/http/middlware"
	"github.com/gorilla/mux"
)

func NewAPIRouter(appCtx *app.AppContext) *mux.Router {
	router := mux.NewRouter()
	// Add middleware
	router.Use(middleware.LoggerMiddleware)
	// Route to create dashboard
	router.HandleFunc("/dashboard", appCtx.DashboardHandler.CreateDashboardHandler).Methods("POST")
	router.HandleFunc("/dashboard", appCtx.DashboardHandler.GetManyDashboardsByIDUserHandler).Methods("GET")
	// router.HandleFunc("/login", )
	return router
}
