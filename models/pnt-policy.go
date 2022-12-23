package models

import (
	"gorm.io/gorm"
	pntPolicy "medici.vn/commission-serivce/enums/pnt-policy"
)

type PntPolicy struct {
	gorm.Model
	Status pntPolicy.Status
}

func (PntPolicy) TableName() string {
	return "pnt_policies"
}
