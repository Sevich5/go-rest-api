package dto

import (
	"app/internal/domain/entity"
	"github.com/google/uuid"
	"time"
)

type UserPublicDto struct {
	Id        uuid.UUID  `json:"id"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at,omitzero"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UserPrivateDto struct {
	Id        uuid.UUID  `json:"id"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at,omitzero"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func NewUserPublicDto(e *entity.User) *UserPublicDto {
	var updatedAt *time.Time
	if !e.UpdatedAt.IsZero() {
		updatedAt = &e.UpdatedAt
	}
	return &UserPublicDto{
		Id:        uuid.MustParse(e.Id.String()),
		Email:     e.Email,
		CreatedAt: e.CreatedAt,
		UpdatedAt: updatedAt,
	}
}

func NewUserPrivateDto(e *entity.User) *UserPrivateDto {
	var updatedAt *time.Time
	if !e.UpdatedAt.IsZero() {
		updatedAt = &e.UpdatedAt
	}
	return &UserPrivateDto{
		Id:        uuid.MustParse(e.Id.String()),
		Email:     e.Email,
		Password:  e.Password,
		CreatedAt: e.CreatedAt,
		UpdatedAt: updatedAt,
	}
}
