package controllers

import (
	"app/internal/presentation/rest/middleware"
	"github.com/gin-gonic/gin"
)

func AddUserGinRouter(userHandler *UserRestHandler, r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("all/", userHandler.GetAllUsers)
		users.POST("create/", middleware.Auth(), userHandler.CreateUser)
	}
}
