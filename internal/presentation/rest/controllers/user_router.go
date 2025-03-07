package controllers

import (
	"app/internal/presentation/rest/middleware/authentication"
	"github.com/gin-gonic/gin"
)

func AddUserGinRouter(userHandler *UserRestHandler, authMiddleware authentication.AuthMiddleware, r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("all/", userHandler.GetAllUsers)
		users.POST("create/", authMiddleware.Auth(), userHandler.CreateUser)
		users.DELETE("delete/", authMiddleware.Auth(), userHandler.DeleteUser)
		users.PATCH("update/", authMiddleware.Auth(), userHandler.UpdateUser)
	}
}
