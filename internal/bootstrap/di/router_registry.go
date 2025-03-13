package di

import (
	"app/internal/presentation/rest/controllers"
	"app/internal/presentation/rest/middleware"
	"app/internal/presentation/rest/middleware/authentication"
	"github.com/gin-gonic/gin"
)

func NewRouterRegistry(services *ServiceRegistry, mode string) *gin.Engine {
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.Use(middleware.RecoveryMiddleware())
	router.NoRoute(middleware.NotFoundMiddleware())
	authMiddleware := authentication.NewJWTAuth(services.TokenToolService)
	userHandler := controllers.NewUserRestHandler(services.UserService)
	authHandler := controllers.NewAuthRestHandler(services.AuthService)
	controllers.AddUserGinRouter(userHandler, authMiddleware, router)
	controllers.AddAuthGinRouter(authHandler, router)
	return router
}
