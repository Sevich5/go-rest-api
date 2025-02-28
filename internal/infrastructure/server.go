package infrastructure

import (
	"app/internal/infrastructure/configuration"
	"app/internal/infrastructure/persistence"
	"app/internal/infrastructure/registry"
)

func ServerRun() {
	cfg := configuration.LoadConfig()
	db := persistence.ConnectDatabase(cfg)
	server := registry.InitWithDbAndConfig(db, cfg)
	err := server.Run(":" + cfg.Server.Port)
	if err != nil {
		panic(err)
	}
}
