package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string `gorm:"type:varchar(255)" json:"full_name"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Token    string `gorm:"-" json:"token,omitempty"`
}
