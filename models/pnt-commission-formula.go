package models

import (
	"gorm.io/gorm"
)

type PntCommissionFormula struct {
	gorm.Model
	PntPolicyId  uint
	LevelCode    string
	PntProductId uint
	Value        float32
}

func (PntCommissionFormula) TableName() string {
	return "pnt_commission_formulas"
}
