package converter

import (
	"app/internal/domain/entity"
	"app/internal/infrastructure/persistence/model"
)

type UserConverter struct{}

func NewUserConverter() *UserConverter {
	return &UserConverter{}
}

func (c *UserConverter) FromDomainToModel(d *entity.User) model.User {
	return model.User{
		Id:        d.Id,
		Email:     d.Email,
		Password:  d.Password,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func (c *UserConverter) FromModelToDomain(m *model.User) entity.User {
	return entity.User{
		Id:        m.Id,
		Email:     m.Email,
		Password:  m.Password,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
