package converter

import (
	"app/internal/domain/entity"
	"app/internal/domain/valueobject"
	"app/internal/infrastructure/persistence/model"
	"time"
)

type UserConverter struct{}

func NewUserConverter() *UserConverter {
	return &UserConverter{}
}

func (c *UserConverter) FromDomainToModel(d *entity.User) *model.User {
	var createdAt *time.Time
	if !d.CreatedAt.IsSet() {
		createdAt = nil
	} else {
		voCreatedAt := d.CreatedAt.Value()
		createdAt = &voCreatedAt
	}
	var updatedAt *time.Time
	if !d.UpdatedAt.IsSet() {
		updatedAt = nil
	} else {
		voUpdatedAt := d.UpdatedAt.Value()
		updatedAt = &voUpdatedAt
	}
	return &model.User{
		Id:        d.Id.Value(),
		Email:     d.Email.Value(),
		Password:  d.Password.Value(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func (c *UserConverter) FromModelToDomain(m *model.User) *entity.User {
	email, _ := valueobject.NewEmail(m.Email)
	password, _ := valueobject.NewPassword(m.Password)
	createdAt := valueobject.NullOptionalTime()
	if m.CreatedAt != nil {
		createdAt = valueobject.NewOptionalTime(*m.CreatedAt)
	}
	updatedAt := valueobject.NullOptionalTime()
	if m.UpdatedAt != nil {
		updatedAt = valueobject.NewOptionalTime(*m.UpdatedAt)
	}
	return &entity.User{
		Id:        valueobject.NewUuidFromUuid(m.Id),
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
