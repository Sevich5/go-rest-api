package controllers

import "github.com/gin-gonic/gin"

func AddAuthGinRouter(userHandler *AuthRestHandler, r *gin.Engine) {
	r.POST("/login", userHandler.Login)
}
