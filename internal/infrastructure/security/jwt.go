package security

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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

	return token.SignedString(t.GetSecret())
}

func NewJWTTokenToolService(tokenKey []byte) *JWTTokenToolService {
	return &JWTTokenToolService{
		AppSecretKey: tokenKey,
	}
}
