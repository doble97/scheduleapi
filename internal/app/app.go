package app

import (
	"os"

	"github.com/doble97/scheduleapi/config"
	"github.com/doble97/scheduleapi/internal/core/ports"
	"github.com/doble97/scheduleapi/internal/core/services"
	"github.com/doble97/scheduleapi/internal/platform/api/http/handler"
	"github.com/doble97/scheduleapi/internal/platform/repository/inmemory"
	"github.com/doble97/scheduleapi/internal/platform/repository/mariadb"
)

type AppContext struct {
	DashboardHandler *handler.DashboardHandler
	UserHandler      *handler.UserHandler
}

func NewAppContext(gC config.GlobalConfig) *AppContext {
	var repo ports.DashboardRepository
	if gC.DBStage == "inmemory" {
		repo = inmemory.NewInMemoryDashboardRepo()
	} else {
		dsn := os.Getenv("DBDsn")
		db, err := mariadb.InitDB(dsn)
		if err != nil {
			panic("Failed to initialize database: " + err.Error())
		}
		repo = mariadb.NewDashboardRepoMariaDB(db)
		// defer db.Close()
	}
	servi := services.NewDashboardService(repo)
	dashboardHandler := handler.NewDashboardHandler(servi)
	userHandler := getUserHandler(gC.DBDsn)
	return &AppContext{
		DashboardHandler: dashboardHandler,
		UserHandler:      &userHandler,
	}
}

func getUserHandler(dsn string) handler.UserHandler {
	var repo ports.UserRepository
	conn, err := mariadb.InitDB(dsn)
	if err != nil {
		panic("Panic to initialize database")
	}
	repo = mariadb.NewUserRepoMariaDB(conn)
	services := services.NewAuthService(repo)
	return *handler.NewUserHandler(services)
}
