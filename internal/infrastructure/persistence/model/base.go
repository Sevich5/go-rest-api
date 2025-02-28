package model

import (
	"github.com/google/uuid"
)

type Base interface {
	GetModelId() uuid.UUID
}
