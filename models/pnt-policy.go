package models

import (
	"gorm.io/gorm"
)

type PntPolicy struct {
	gorm.Model
	Status string
}

func (PntPolicy) TableName() string {
	return "pnt_policies"
}
