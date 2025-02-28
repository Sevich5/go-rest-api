package repository

import (
	"app/internal/application/dto"
	"app/internal/domain/entity"
	"app/internal/domain/interfaces"
	"app/internal/infrastructure/persistence/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db        *gorm.DB
	converter *dto.UserConverter
}

func NewUserPgRepository(db *gorm.DB) interfaces.UserRepository {
	repository := &UserRepository{
		db:        db,
		converter: dto.NewUserConverter(),
	}
	return repository
}

func (r *UserRepository) GetAll() ([]*entity.User, error) {
	var items []model.User
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}
	users := make([]*entity.User, len(items))
	for i, item := range items {
		converted := r.converter.FromModelToDomain(&item)
		users[i] = &converted
	}
	return users, nil
}

func (r *UserRepository) Create(user *entity.User) error {
	userModel := r.converter.FromDomainToModel(user)
	if err := r.db.Create(&userModel).Error; err != nil {
		return err
	}
	user.Id = userModel.GetModelId()
	return nil
}

func (r *UserRepository) GetById(uuid uuid.UUID) (entity.User, error) {
	userModel := &model.User{}
	if err := r.db.Where("id = ?", uuid).First(&userModel).Error; err != nil {
		return entity.User{}, err
	}
	return r.converter.FromModelToDomain(userModel), nil
}

func (r *UserRepository) GetByEmail(email string) (entity.User, error) {
	userModel := &model.User{}
	if err := r.db.Where("email = ?", email).First(&userModel).Error; err != nil {
		return entity.User{}, err
	}
	return r.converter.FromModelToDomain(userModel), nil
}

func (r *UserRepository) Update(user *entity.User) error {
	userModel := r.converter.FromDomainToModel(user)
	if err := r.db.Save(&userModel).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Delete(user *entity.User) error {
	userModel := r.converter.FromDomainToModel(user)
	if err := r.db.Where("id = ?", userModel.GetModelId()).Delete(&userModel).Error; err != nil {
		return err
	}
	return nil
}
