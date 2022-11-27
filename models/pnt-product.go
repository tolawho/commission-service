package models

import (
	"gorm.io/gorm"
)

type PntProduct struct {
	gorm.Model
	Code               string
	PntCategoryId      uint
	InsuranceCompanyId uint
	Name               string
	Status             int
	CommissionRate     float32
	Amount             float32
	Vat                float32
}

func (PntProduct) TableName() string {
	return "pnt_products"
}
