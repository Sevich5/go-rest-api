package requestdto

import "github.com/google/uuid"

type CreateUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDto struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
