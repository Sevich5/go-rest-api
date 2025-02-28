package interfaces

import (
	"app/internal/domain/entity"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetAll() ([]*entity.User, error)
	Create(user *entity.User) error
	GetById(uuid uuid.UUID) (entity.User, error)
	GetByEmail(email string) (entity.User, error)
	Update(user *entity.User) error
	Delete(user *entity.User) error
}
