package models

import (
	"gorm.io/gorm"
)

type PntContractProduct struct {
	gorm.Model
	PntContractId  uint
	PntProductId   uint
	Group          string
	Amount         float32
	ExtraAmount    float32
	Tax            float32
	CommissionRate float32
}

func (PntContractProduct) TableName() string {
	return "pnt_contract_product"
}
