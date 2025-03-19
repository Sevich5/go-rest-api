package repository

import (
	"app/internal/application"
	"app/internal/domain/entity"
	"app/internal/domain/interfaces"
	"app/internal/infrastructure/persistence/converter"
	"app/internal/infrastructure/persistence/model"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db        *gorm.DB
	converter *converter.UserConverter
}

func NewUserPgRepository(db *gorm.DB) interfaces.UserRepository {
	return &UserRepository{
		db:        db,
		converter: converter.NewUserConverter(),
	}
}

func (r *UserRepository) GetAll(limit int, offset int) (users []*entity.User, err error, limitOut int, offsetOut int) {
	var items []model.User
	if err := r.db.Limit(limit).Offset(offset).Find(&items).Error; err != nil {
		return nil, application.NewErrorFromErr(err), 0, 0
	}
	users = make([]*entity.User, len(items))
	for i, item := range items {
		users[i] = r.converter.FromModelToDomain(&item)
	}
	return users, nil, limit, offset
}

func (r *UserRepository) Create(user *entity.User) error {
	userModel := r.converter.FromDomainToModel(user)
	if err := r.db.Create(&userModel).Error; err != nil {
		return application.NewErrorFromErr(err)
	}
	return nil
}

func (r *UserRepository) GetById(uuid uuid.UUID) (*entity.User, error) {
	userModel := &model.User{}
	err := r.db.Where("id = ?", uuid).First(&userModel).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, application.NewErrorFromErr(err)
	}
	return r.converter.FromModelToDomain(userModel), nil
}

func (r *UserRepository) GetByEmail(email string) (*entity.User, error) {
	userModel := &model.User{}
	err := r.db.Where("email = ?", email).First(&userModel).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, application.NewErrorFromErr(err)
	}
	return r.converter.FromModelToDomain(userModel), nil
}

func (r *UserRepository) Update(user *entity.User) error {
	userModel := r.converter.FromDomainToModel(user)
	if err := r.db.Save(&userModel).Error; err != nil {
		return application.NewErrorFromErr(err)
	}
	return nil
}

func (r *UserRepository) Delete(user *entity.User) error {
	userModel := r.converter.FromDomainToModel(user)
	if err := r.db.Where("id = ?", userModel.GetModelId()).Delete(&userModel).Error; err != nil {
		return application.NewErrorFromErr(err)
	}
	return nil
}
