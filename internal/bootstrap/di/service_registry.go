package di

import (
	"app/internal/application"
	"app/internal/infrastructure/configuration"
	"app/internal/infrastructure/repository"
	"app/internal/infrastructure/security"
	"gorm.io/gorm"
)

type ServiceRegistry struct {
	TokenToolService *security.JWTTokenToolService
	AuthService      *application.AuthService
	UserService      *application.UserService
}

func NewServiceRegistry(db *gorm.DB, cfg *configuration.Config) *ServiceRegistry {
	userRepository := repository.NewUserPgRepository(db)
	tokenToolService := security.NewJWTTokenToolService(cfg.Application.SecretKey)
	return &ServiceRegistry{
		TokenToolService: tokenToolService,
		AuthService:      application.NewAuthService(userRepository, tokenToolService),
		UserService:      application.NewUserService(userRepository),
	}
}
