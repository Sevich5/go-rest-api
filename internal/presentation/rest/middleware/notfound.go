package middleware

import (
	"app/internal/presentation/helpers"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotFoundMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		helpers.JsonError(c, errors.New("page not found"), http.StatusNotFound)
	}
}
