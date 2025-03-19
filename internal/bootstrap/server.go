package bootstrap

import (
	"app/internal/bootstrap/di"
	"app/internal/infrastructure/configuration"
	"app/internal/infrastructure/persistence"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ServerRun() {
	cfg := configuration.LoadConfig()
	db := persistence.ConnectDatabase(cfg)
	services := di.NewServiceRegistry(db, cfg)
	router := di.NewRouterRegistry(services, cfg.Application.Mode)
	server := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: router.Handler(),
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Listen: %s\n", err)
		}
	}()
	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-quitSignal
	log.Println("Gracefully shutting down server after 6 seconds...")
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server shutdown:", err)
	}
	<-ctx.Done()
}
