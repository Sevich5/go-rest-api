package controllers

import (
	"app/internal/application"
	"app/internal/application/dto"
	"app/internal/application/requestdto"
	"app/internal/presentation/helpers"
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
	paginatedDto := requestdto.PaginatedDto{}
	err := c.ShouldBindQuery(&paginatedDto)
	if err != nil {
		helpers.JsonError(c, err, http.StatusBadRequest)
		return
	}
	users, err, limit, offset := h.service.GetAllUsers(paginatedDto)
	if err != nil {
		helpers.JsonError(c, err, http.StatusBadRequest)
		return
	}
	var dtoUsers []interface{}
	for _, user := range users {
		dtoUsers = append(dtoUsers, dto.NewUserPublicDto(user))
	}
	helpers.JsonList(c, dtoUsers, limit, offset, len(dtoUsers))
	return
}

func (h *UserRestHandler) CreateUser(c *gin.Context) {
	userJson := requestdto.UserDto{}
	err := c.BindJSON(&userJson)
	if err != nil {
		helpers.JsonError(c, err, http.StatusBadRequest)
		return
	}
	user, err := h.service.CreateUser(userJson.Email, userJson.Password)
	if err != nil {
		helpers.JsonError(c, err, http.StatusBadRequest)
		return
	}
	helpers.JsonOk(c, dto.NewUserPrivateDto(user))
	return
}

func (h *UserRestHandler) UpdateUser(c *gin.Context) {
	userJson := requestdto.UserDto{}
	err := c.ShouldBindJSON(&userJson)
	if err != nil {
		helpers.JsonError(c, err, http.StatusBadRequest)
		return
	}
	user, err := h.service.UpdateUser(userJson)
	if err != nil {
		helpers.JsonError(c, err, http.StatusBadRequest)
		return
	}
	helpers.JsonOk(c, dto.NewUserPublicDto(user))
	return
}

func (h *UserRestHandler) DeleteUser(c *gin.Context) {
	userJson := requestdto.UserDto{}
	err := c.ShouldBindJSON(&userJson)
	if err != nil {
		helpers.JsonError(c, err, http.StatusBadRequest)
		return
	}
	err = h.service.DeleteUser(userJson.Id.String())
	if err != nil {
		helpers.JsonError(c, err, http.StatusBadRequest)
		return
	}
	helpers.JsonOk(c, nil)
	return
}
