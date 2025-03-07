package authentication

import (
	"app/internal/application"
	"app/internal/presentation/helpers"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

type JWTAuth struct {
	tool application.TokenTool
}

func (j *JWTAuth) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := j.extractToken(c)
		if tokenString == "" {
			c.Abort()
			helpers.JsonError(c, errors.New("no token"), http.StatusUnauthorized)
			return
		}
		claims, err := j.ValidateJWT(tokenString, j.tool.GetSecret())
		if err != nil {
			c.Abort()
			helpers.JsonError(c, err, http.StatusUnauthorized)
			return
		}
		c.Set("user_id", claims["user_id"])
		c.Set("email", claims["email"])

		c.Next()
	}
}

func (j *JWTAuth) extractToken(c *gin.Context) string {
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

func (j *JWTAuth) ValidateJWT(tokenString string, secret []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token is invalid")
}

func NewJWTAuth(tool application.TokenTool) *JWTAuth {
	return &JWTAuth{tool: tool}
}
