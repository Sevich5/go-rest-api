package registry

import (
	"app/internal/application"
	"app/internal/infrastructure/configuration"
	"app/internal/infrastructure/repository"
	"app/internal/presentation/rest/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitWithDbAndConfig(db *gorm.DB, cfg *configuration.Config) *gin.Engine {
	if cfg.Application.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	userRepository := repository.NewUserPgRepository(db)
	UserService := application.NewUserService(userRepository)
	userHandler := controllers.NewUserRestHandler(UserService)
	authHandler := controllers.NewAuthRestHandler(UserService)
	controllers.AddUserGinRouter(userHandler, router)
	controllers.AddAuthGinRouter(authHandler, router)
	return router
}
