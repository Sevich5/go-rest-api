package security

import (
	"app/internal/infrastructure/configuration"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

func GetJWTSecret() []byte {
	return configuration.AppSecretKey
}

// GenerateJWT Создание JWT-токена
func GenerateJWT(userID uuid.UUID, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24 * 365 * 5).Unix(),
	})

	return token.SignedString(GetJWTSecret())
}
