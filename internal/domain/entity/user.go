package entity

import (
	"app/internal/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Root
	Id        uuid.UUID
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(email string, password string) (*User, error) {
	if email == "" {
		return nil, domain.NewError("email is required")
	}
	if password == "" {
		return nil, domain.NewError("password is required")
	}
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, domain.NewError(err.Error())
	}
	//TODO need validation in ValuesObject
	return &User{
		Id:        uuid.New(),
		Email:     email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", domain.NewError(err.Error())
	}
	return string(hashedPassword), nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) GetIdString() string {
	return u.Id.String()
}

func (u *User) SetUpdatedAt() {
	u.UpdatedAt = time.Now()
}
