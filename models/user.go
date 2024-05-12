package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
