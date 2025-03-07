package bootstrap

import (
	"app/internal/bootstrap/di"
	"app/internal/infrastructure/configuration"
	"app/internal/infrastructure/persistence"
)

func ServerRun() {
	cfg := configuration.LoadConfig()
	db := persistence.ConnectDatabase(cfg)
	services := di.NewServiceRegistry(db, cfg)
	server := di.NewRouterRegistry(services, cfg.Application.Mode)
	err := server.Run(":" + cfg.Server.Port)
	if err != nil {
		panic(err)
	}
}
