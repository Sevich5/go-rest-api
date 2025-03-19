package valueobject

import (
	"app/internal/domain"
	"regexp"
)

type Email struct {
	value string
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func NewEmail(email string) (Email, error) {
	if len(email) == 0 || !emailRegex.MatchString(email) {
		return Email{}, domain.NewError("Invalid email format")
	}
	return Email{value: email}, nil
}

func (e *Email) Value() string {
	return e.value
}

type Password struct {
	value string
}

func NewPassword(value string) (Password, error) {
	if len(value) < 8 {
		return Password{}, domain.NewError("Password must be at least 8 characters")
	}
	return Password{value: value}, nil
}

func (p *Password) Value() string {
	return p.value
}

func (p *Password) SetValue(value string) {
	p.value = value
}
