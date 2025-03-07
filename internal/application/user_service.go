package application

import (
	"app/internal/application/requestdto"
	"app/internal/domain/entity"
	"app/internal/domain/interfaces"
	"errors"
	"github.com/google/uuid"
)

type UserService struct {
	UserRepository interfaces.UserRepository
	TokenTool      TokenTool
}

func NewUserService(userRepository interfaces.UserRepository, tokenTool TokenTool) *UserService {
	return &UserService{
		UserRepository: userRepository,
		TokenTool:      tokenTool,
	}
}

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.UserRepository.GetByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}
	if err := user.CheckPassword(password); err != nil {
		return "", errors.New("wrong password")
	}
	token, err := s.TokenTool.GenerateToken(user.Id, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) CreateUser(email string, password string) (*entity.User, error) {
	found, _ := s.UserRepository.GetByEmail(email)
	if found.GetIdString() != uuid.Nil.String() {
		return nil, errors.New("user with this email already exists")
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
	user.SetUpdatedAt()
	return nil
}

func (s *UserService) DeleteUser(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	err = s.UserRepository.Delete(uid)
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

func (s *UserService) GetAllUsers(dto requestdto.PaginatedDto) (users []*entity.User, err error, limitOut int, offsetOut int) {
	users, err, limitOut, offsetOut = s.UserRepository.GetAll(dto.Limit, dto.Offset)
	if err != nil {
		return nil, err, 0, 0
	}
	return users, nil, limitOut, offsetOut
}
