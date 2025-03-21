package controllers

import (
	"app/internal/application"
	"app/internal/application/requestdto"
	"app/internal/presentation/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthRestHandler struct {
	service *application.AuthService
}

func NewAuthRestHandler(service *application.AuthService) *AuthRestHandler {
	return &AuthRestHandler{service: service}
}

func (h *AuthRestHandler) Login(c *gin.Context) {
	var requestData requestdto.LoginRequest
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		helpers.JsonError(c, err, http.StatusBadRequest)
		return
	}
	email := requestData.Email
	password := requestData.Password
	token, err := h.service.Login(email, password)
	if err != nil {
		helpers.JsonError(c, err, http.StatusBadRequest)
		return
	}
	helpers.JsonOk(c, token)
	return
}
