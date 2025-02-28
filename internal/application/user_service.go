package application

import (
	"app/internal/domain/entity"
	"app/internal/domain/interfaces"
	"app/internal/infrastructure/security"
	"errors"
	"github.com/google/uuid"
)

type UserService struct {
	UserRepository interfaces.UserRepository
}

func NewUserService(userRepository interfaces.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) Login(email, password string) (string, error) {
	// Ищем пользователя в репозитории
	user, err := s.UserRepository.GetByEmail(email)
	if err != nil {
		return "", errors.New("пользователь не найден")
	}

	// Проверяем пароль
	if err := user.CheckPassword(password); err != nil {
		return "", errors.New("неверный пароль")
	}

	// Генерируем JWT-токен
	token, err := security.GenerateJWT(user.Id, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) CreateUser(email string, password string) (*entity.User, error) {
	found, _ := s.UserRepository.GetByEmail(email)
	if found.GetIdString() != "" {
		return nil, errors.New("пользователь с таким email уже существует")
	}
	user, err := entity.NewUser(email, password)
	if err != nil {
		return nil, err
	}
	err = s.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(user *entity.User, newPassword string) error {
	if newPassword != "" {
		hashedPassword, err := entity.HashPassword(newPassword)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	}
	err := s.UserRepository.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(user *entity.User) error {
	err := s.UserRepository.Delete(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUserById(id string) (*entity.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	user, err := s.UserRepository.GetById(uid)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetUserByEmail(email string) (*entity.User, error) {
	user, err := s.UserRepository.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetAllUsers() ([]*entity.User, error) {
	users, err := s.UserRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
