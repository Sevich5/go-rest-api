package middleware

import (
	"app/internal/infrastructure/security"
	"app/internal/presentation/helper"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := extractToken(c)
		if tokenString == "" {
			c.Abort()
			helper.JsonError(c, errors.New("токен отсутствует"), http.StatusUnauthorized)
		}

		claims, err := ValidateJWT(tokenString)
		if err != nil {
			c.Abort()
			helper.JsonError(c, err, http.StatusUnauthorized)
		}

		// Добавляем данные пользователя в контекст запроса
		c.Set("user_id", claims["user_id"])
		c.Set("email", claims["email"])

		c.Next() // Передаем управление дальше
	}
}

// Извлечение токена из заголовка Authorization
func extractToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}

// ValidateJWT Валидация JWT-токена
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return security.GetJWTSecret(), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("недействительный токен доступа")
}
