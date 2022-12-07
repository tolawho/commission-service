package models

import (
	"gorm.io/gorm"
)

type PntContract struct {
	gorm.Model
	ID                  uint
	AgencyId            uint
	PntCategoryId       uint
	InsuranceCompanyId  uint
	Status              int
	ProviderStatus      int
	GcnStatus           int
	Type                int
	Amount              float32
	CommissionRate      float32
	Agency              *Agency               `gorm:"foreignKey:AgencyId"`
	PntContractProducts []*PntContractProduct `gorm:"foreignKey:PntContractId"`
}

func (PntContract) TableName() string {
	return "pnt_contracts"
}
