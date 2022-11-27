package models

import (
	"gorm.io/gorm"
)

type PntAgencyTree struct {
	gorm.Model
	AgencyId uint
	ParentId uint
}

func (PntAgencyTree) TableName() string {
	return "pnt_agency_tree"
}
