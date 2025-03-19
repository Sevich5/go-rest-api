package application

import (
	"app/internal/application/requestdto"
	"app/internal/domain/entity"
	"app/internal/domain/interfaces"
	"app/internal/domain/valueobject"
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
	if found.Id.Value() != uuid.Nil {
		appErr := NewError("User with this email already exists")
		appErr.StatusCode = http.StatusConflict
		return nil, appErr
	}
	voEmail, err := valueobject.NewEmail(email)
	if err != nil {
		return nil, NewErrorFromErr(err)
	}
	voPassword, err := valueobject.NewPassword(password)
	if err != nil {
		return nil, NewErrorFromErr(err)
	}
	user, err := entity.NewUser(voEmail, voPassword)
	if err != nil {
		return nil, NewErrorFromErr(err)
	}
	if err = s.UserRepository.Create(user); err != nil {
		return nil, NewErrorFromErr(err)
	}
	return user, nil
}

func (s *UserService) UpdateUser(user *entity.User, dto requestdto.UserDto) error {
	if dto.Password != "" {
		if _, err := valueobject.NewPassword(dto.Password); err != nil {
			return NewErrorFromErr(err)
		}
		hashedPassword, err := entity.HashPassword(dto.Password)
		if err != nil {
			return NewErrorFromErr(err)
		}
		user.Password, _ = valueobject.NewPassword(hashedPassword)
	}
	if dto.Email != "" {
		voEmail, err := valueobject.NewEmail(dto.Email)
		if err != nil {
			return NewErrorFromErr(err)
		}
		user.Email = voEmail
	}
	err := s.UserRepository.Update(user)
	if err != nil {
		return NewErrorFromErr(err)
	}
	user.SetUpdatedNow()
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
	return user, nil
}

func (s *UserService) GetUserByEmail(email string) (*entity.User, error) {
	user, err := s.UserRepository.GetByEmail(email)
	if err != nil {
		return nil, NewErrorFromErr(err)
	}
	return user, nil
}

func (s *UserService) GetAllUsers(dto requestdto.PaginatedDto) (users []*entity.User, err error, limitOut int, offsetOut int) {
	users, err, limitOut, offsetOut = s.UserRepository.GetAll(dto.Limit, dto.Offset)
	if err != nil {
		return nil, NewErrorFromErr(err), 0, 0
	}
	return users, nil, limitOut, offsetOut
}
