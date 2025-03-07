package interfaces

import (
	"app/internal/domain/entity"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetAll(limit int, offset int) (users []*entity.User, error error, limitOut int, offsetOut int)
	Create(user *entity.User) error
	GetById(uuid uuid.UUID) (entity.User, error)
	GetByEmail(email string) (entity.User, error)
	Update(user *entity.User) error
	Delete(uuid uuid.UUID) error
}
