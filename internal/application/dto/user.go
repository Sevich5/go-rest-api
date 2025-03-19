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
	if e.UpdatedAt.IsSet() {
		voUpdatedAt := e.UpdatedAt.Value()
		updatedAt = &voUpdatedAt
	}
	return &UserPublicDto{
		Id:        e.Id.Value(),
		Email:     e.Email.Value(),
		CreatedAt: e.CreatedAt.Value(),
		UpdatedAt: updatedAt,
	}
}

func NewUserPrivateDto(e *entity.User) *UserPrivateDto {
	var updatedAt *time.Time
	if e.UpdatedAt.IsSet() {
		voUpdatedAt := e.UpdatedAt.Value()
		updatedAt = &voUpdatedAt
	}
	return &UserPrivateDto{
		Id:        e.Id.Value(),
		Email:     e.Email.Value(),
		Password:  e.Password.Value(),
		CreatedAt: e.CreatedAt.Value(),
		UpdatedAt: updatedAt,
	}
}
