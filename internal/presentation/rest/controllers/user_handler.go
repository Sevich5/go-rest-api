package controllers

import (
	"app/internal/application"
	"app/internal/application/dto"
	"app/internal/infrastructure/persistence/model"
	"app/internal/presentation/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserRestHandler struct {
	service *application.UserService
}

func NewUserRestHandler(service *application.UserService) *UserRestHandler {
	return &UserRestHandler{service: service}
}

func (h *UserRestHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		helper.JsonError(c, err, http.StatusBadRequest)
		return
	}
	var dtoUsers []interface{}
	for _, user := range users {
		dtoUsers = append(dtoUsers, dto.UserPublicDto(user))
	}
	helper.JsonList(c, dtoUsers, 0, 0, len(dtoUsers))
	return
}

func (h *UserRestHandler) CreateUser(c *gin.Context) {
	userJson := model.User{}
	err := c.BindJSON(&userJson)
	if err != nil {
		helper.JsonError(c, err, http.StatusBadRequest)
		return
	}
	user, err := h.service.CreateUser(userJson.Email, userJson.Password)
	if err != nil {
		helper.JsonError(c, err, http.StatusBadRequest)
		return
	}
	helper.JsonOk(c, dto.UserPublicDto(user))
	return
}
