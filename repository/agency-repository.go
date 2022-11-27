package repository

import (
	"gorm.io/gorm"
	"medici.vn/commission-serivce/models"
)

// AgencyRepository  is contract what agencyRepository can do to db
type AgencyRepository interface {
	FindById(Id uint) *models.Agency
}

type agencyConnection struct {
	connection *gorm.DB
}

func (db agencyConnection) FindById(Id uint) *models.Agency {
	var agency *models.Agency
	db.connection.Where("id = ?", Id).First(&agency)
	return agency
}

// NewAgencyRepository is creates a new instance of AgencyRepository
func NewAgencyRepository(db *gorm.DB) AgencyRepository {
	return &agencyConnection{
		connection: db,
	}
}
