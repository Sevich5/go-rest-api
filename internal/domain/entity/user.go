package entity

import (
	"app/internal/domain"
	"app/internal/domain/valueobject"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	AggregateRoot
	Id        valueobject.Uuid
	Email     valueobject.Email
	Password  valueobject.Password
	CreatedAt valueobject.OptionalTime
	UpdatedAt valueobject.OptionalTime
}

type UserCreatedEvent struct {
	DomainEventInterface
	UserID valueobject.Uuid
	Email  valueobject.Email
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", domain.NewError(err.Error())
	}
	return string(hashedPassword), nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password.Value()), []byte(password))
}

func (u *User) SetUpdatedNow() {
	u.UpdatedAt = valueobject.NewOptionalTime(time.Now())
}

func NewUser(email valueobject.Email, password valueobject.Password) (*User, error) {
	hashedPassword := valueobject.Password{}
	if password.Value() != "" {
		hashed, err := HashPassword(password.Value())
		if err != nil {
			return nil, err
		}
		hashedPassword, err = valueobject.NewPassword(hashed)
		if err != nil {
			return nil, err
		}
	}
	user := &User{
		AggregateRoot: *NewAggregateRoot(),
		Id:            valueobject.NewUuid(),
		Email:         email,
		Password:      hashedPassword,
	}
	user.CreatedAt = valueobject.NewOptionalTime(time.Now())
	user.AggregateRoot.AddDomainEvent(&UserCreatedEvent{UserID: user.Id, Email: user.Email})
	return user, nil
}
