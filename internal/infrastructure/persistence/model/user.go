package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Base      `gorm:"-"`
	Id        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Email     string    `gorm:"uniqueIndex"`
	Password  string    `json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"default:null"`
}

func (u *User) GetModelId() uuid.UUID {
	return u.Id
}
