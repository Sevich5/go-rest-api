package middleware

import (
	"app/internal/presentation/helpers"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v", err)
				helpers.JsonError(c, errors.New(fmt.Sprintf("%s", err)), http.StatusInternalServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
