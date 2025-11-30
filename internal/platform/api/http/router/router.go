// Package router implements the API routing for the application
package router

import (
	"net/http"

	"github.com/doble97/scheduleapi/internal/app"
	"github.com/doble97/scheduleapi/internal/platform/api/http/handler"
	middleware "github.com/doble97/scheduleapi/internal/platform/api/http/middlware"
	"github.com/doble97/scheduleapi/internal/platform/api/http/util"
	"github.com/gorilla/mux"
)

func NewAPIRouter(appCtx *app.AppContext) *mux.Router {
	router := mux.NewRouter()
	// Add middleware
	router.Use(middleware.LoggerMiddleware)
	// Route to create dashboard
	authMiddleware := middleware.AuthMiddlewareFactory(util.VerifyJWT)
	dashRoute := router.PathPrefix("/dashboard").Subrouter()
	dashRoute.Use(authMiddleware)
	dashRoute.HandleFunc("", handler.ErrorHandlerMiddleware(appCtx.DashboardHandler.CreateDashboardHandler)).Methods("POST")
	dashRoute.HandleFunc("/{userID}", handler.ErrorHandlerMiddleware(appCtx.DashboardHandler.GetManyDashboardsByIDUserHandler)).Methods("GET")
	fs := http.FileServer(http.Dir("./tmp"))
	router.PathPrefix("/files/").Handler(http.StripPrefix("/files/", fs))
	router.HandleFunc("/register", handler.ErrorHandlerMiddleware(appCtx.UserHandler.RegisterHandler)).Methods("POST")
	router.HandleFunc("/login", handler.ErrorHandlerMiddleware(appCtx.UserHandler.LoginHandler)).Methods("POST")
	return router
}
