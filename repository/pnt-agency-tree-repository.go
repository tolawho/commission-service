package repository

import (
	"gorm.io/gorm"
	"medici.vn/commission-serivce/models"
)

// PntAgencyTreeRepository  is contract what pntAgencyTreeRepository can do to db
type PntAgencyTreeRepository interface {
	FindByAgencyId(AgencyId uint) *models.PntAgencyTree
}

type pntAgencyTreeConnection struct {
	connection *gorm.DB
}

func (db pntAgencyTreeConnection) FindByAgencyId(AgencyId uint) *models.PntAgencyTree {
	var pntAgencyTree *models.PntAgencyTree
	db.connection.Where("agency_id = ?", AgencyId).First(&pntAgencyTree)
	return pntAgencyTree
}

// NewPntAgencyTreeRepository is creates a new instance of PntAgencyTreeRepository
func NewPntAgencyTreeRepository(db *gorm.DB) PntAgencyTreeRepository {
	return &pntAgencyTreeConnection{
		connection: db,
	}
}
