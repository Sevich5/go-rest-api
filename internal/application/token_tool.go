package application

import (
	"github.com/google/uuid"
)

type TokenTool interface {
	GetSecret() []byte
	GenerateToken(userID uuid.UUID, email string) (string, error)
}
