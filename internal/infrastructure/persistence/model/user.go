package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Base      `gorm:"-"`
	Id        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (u *User) GetModelId() uuid.UUID {
	return u.Id
}
