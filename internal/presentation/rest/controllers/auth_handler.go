package controllers

import (
	"app/internal/application"
	"app/internal/presentation/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthRestHandler struct {
	service *application.UserService
}

func NewAuthRestHandler(service *application.UserService) *AuthRestHandler {
	return &AuthRestHandler{service: service}
}

func (h *AuthRestHandler) Login(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.BindJSON(&requestData)
	if err != nil {
		helpers.JsonError(c, err, http.StatusBadRequest)
		return
	}
	email := requestData["email"].(string)
	password := requestData["password"].(string)
	token, err := h.service.Login(email, password)
	if err != nil {
		helpers.JsonError(c, err, http.StatusBadRequest)
		return
	}
	helpers.JsonOk(c, token)
	return
}
