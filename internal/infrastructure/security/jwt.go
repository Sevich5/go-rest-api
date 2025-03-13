package security

import (
	"app/internal/application"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type JWTTokenToolService struct {
	AppSecretKey []byte
}

func (t *JWTTokenToolService) GetSecret() []byte {
	return t.AppSecretKey
}

func (t *JWTTokenToolService) GenerateToken(userID uuid.UUID, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24 * 365 * 5).Unix(),
	})

	tokenString, err := token.SignedString(t.AppSecretKey)
	if err != nil {
		appErr := application.NewError(err.Error())
		appErr.StatusCode = http.StatusServiceUnavailable
		return "", appErr
	}

	return tokenString, nil
}

func NewJWTTokenToolService(tokenKey []byte) *JWTTokenToolService {
	return &JWTTokenToolService{
		AppSecretKey: tokenKey,
	}
}
