package application

import (
	"app/internal/application/requestdto"
	"app/internal/domain/entity"
	"app/internal/domain/interfaces"
	"github.com/google/uuid"
	"net/http"
)

type UserService struct {
	UserRepository interfaces.UserRepository
}

func NewUserService(userRepository interfaces.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) CreateUser(email string, password string) (*entity.User, error) {
	found, _ := s.UserRepository.GetByEmail(email)
	if found.GetIdString() != uuid.Nil.String() {
		appErr := NewError("user with this email already exists")
		appErr.StatusCode = http.StatusConflict
		return nil, appErr
	}
	user, err := entity.NewUser(email, password)
	if err != nil {
		return nil, NewErrorFromErr(err)
	}
	err = s.UserRepository.Create(user)
	if err != nil {
		return nil, NewErrorFromErr(err)
	}
	return user, nil
}

func (s *UserService) UpdateUser(user *entity.User, newPassword string) error {
	if newPassword != "" {
		hashedPassword, err := entity.HashPassword(newPassword)
		if err != nil {
			return NewErrorFromErr(err)
		}
		user.Password = hashedPassword
	}
	err := s.UserRepository.Update(user)
	if err != nil {
		return NewErrorFromErr(err)
	}
	user.SetUpdatedAt()
	return nil
}

func (s *UserService) DeleteUser(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return NewErrorFromErr(err)
	}
	err = s.UserRepository.Delete(uid)
	if err != nil {
		return NewErrorFromErr(err)
	}
	return nil
}

func (s *UserService) GetUserById(id string) (*entity.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, NewErrorFromErr(err)
	}
	user, err := s.UserRepository.GetById(uid)
	if err != nil {
		appError := NewError(err.Error())
		appError.StatusCode = http.StatusNotFound
		return nil, appError
	}
	return &user, nil
}

func (s *UserService) GetUserByEmail(email string) (*entity.User, error) {
	user, err := s.UserRepository.GetByEmail(email)
	if err != nil {
		return nil, NewErrorFromErr(err)
	}
	return &user, nil
}

func (s *UserService) GetAllUsers(dto requestdto.PaginatedDto) (users []*entity.User, err error, limitOut int, offsetOut int) {
	users, err, limitOut, offsetOut = s.UserRepository.GetAll(dto.Limit, dto.Offset)
	if err != nil {
		return nil, NewErrorFromErr(err), 0, 0
	}
	return users, nil, limitOut, offsetOut
}
