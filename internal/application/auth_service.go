package application

import (
	"app/internal/domain/interfaces"
	"net/http"
)

type AuthService struct {
	UserRepository interfaces.UserRepository
	TokenTool      TokenTool
}

func NewAuthService(userRepository interfaces.UserRepository, tokenTool TokenTool) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
		TokenTool:      tokenTool,
	}
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.UserRepository.GetByEmail(email)
	if err != nil {
		return "", NewAppErrorWithStatus("User not found", http.StatusNotFound)
	}
	if err := user.CheckPassword(password); err != nil {
		return "", NewAppErrorWithStatus("Wrong password", http.StatusUnauthorized)
	}
	token, err := s.TokenTool.GenerateToken(user.Id.Value(), user.Email.Value())
	if err != nil {
		return "", err
	}

	return token, nil
}
