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
	if found != nil && found.Id.Value() != uuid.Nil {
		return nil, NewAppErrorWithStatus("User with this email already exists", http.StatusConflict)
	}
	voEmail, err := valueobject.NewEmail(email)
	if err != nil {
		return nil, NewAppErrorFromErr(err)
	}
	voPassword, err := valueobject.NewPassword(password)
	if err != nil {
		return nil, NewAppErrorFromErr(err)
	}
	user, err := entity.NewUser(voEmail, voPassword)
	if err != nil {
		return nil, NewAppErrorFromErr(err)
	}
	if err = s.UserRepository.Create(user); err != nil {
		return nil, NewAppErrorFromErr(err)
	}
	return user, nil
}

func (s *UserService) UpdateUser(dto requestdto.UserDto) (*entity.User, error) {
	user, err := s.GetUserById(dto.Id.String())
	if err != nil {
		return nil, NewAppErrorFromErr(err)
	}
	if user == nil {
		return nil, NewAppErrorWithStatus("User not found", http.StatusNotFound)
	}
	if dto.Password != "" {
		if _, err := valueobject.NewPassword(dto.Password); err != nil {
			return nil, NewAppErrorFromErr(err)
		}
		hashedPassword, err := entity.HashPassword(dto.Password)
		if err != nil {
			return nil, NewAppErrorFromErr(err)
		}
		user.Password, _ = valueobject.NewPassword(hashedPassword)
	}
	if dto.Email != "" {
		voEmail, err := valueobject.NewEmail(dto.Email)
		if err != nil {
			return nil, NewAppErrorFromErr(err)
		}
		user.Email = voEmail
	}
	err = s.UserRepository.Update(user)
	if err != nil {
		return nil, NewAppErrorFromErr(err)
	}
	user.SetUpdatedNow()
	return user, nil
}

func (s *UserService) DeleteUser(id string) error {
	user, err := s.GetUserById(id)
	if err != nil {
		return NewAppErrorFromErr(err)
	}
	if user == nil {
		return NewAppErrorWithStatus("User not found", http.StatusNotFound)
	}
	err = s.UserRepository.Delete(user)
	if err != nil {
		return NewAppErrorFromErr(err)
	}
	return nil
}

func (s *UserService) GetUserById(id string) (*entity.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, NewAppErrorFromErr(err)
	}
	user, err := s.UserRepository.GetById(uid)
	if err != nil {
		return nil, NewAppErrorWithStatus(err.Error(), http.StatusNotFound)
	}
	if user == nil {
		return nil, NewAppErrorWithStatus("User not found", http.StatusNotFound)
	}
	return user, nil
}

func (s *UserService) GetUserByEmail(email string) (*entity.User, error) {
	user, err := s.UserRepository.GetByEmail(email)
	if err != nil {
		return nil, NewAppErrorFromErr(err)
	}
	return user, nil
}

func (s *UserService) GetAllUsers(dto requestdto.PaginatedDto) (users []*entity.User, err error, limitOut int, offsetOut int) {
	users, err, limitOut, offsetOut = s.UserRepository.GetAll(dto.Limit, dto.Offset)
	if err != nil {
		return nil, NewAppErrorFromErr(err), 0, 0
	}
	return users, nil, limitOut, offsetOut
}
