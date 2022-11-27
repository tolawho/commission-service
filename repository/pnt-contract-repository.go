package repository

import (
	"gorm.io/gorm"
	"medici.vn/commission-serivce/models"
)

// PntContractRepository  is contract what pntContractRepository can do to db
type PntContractRepository interface {
	FindById(Id uint) models.PntContract
}

type pntContractConnection struct {
	connection *gorm.DB
}

func (db pntContractConnection) FindById(Id uint) models.PntContract {
	var pntContract models.PntContract
	db.connection.Where("id = ?", Id).First(&pntContract)
	return pntContract
}

// NewPntContractRepository is creates a new instance of PntContractRepository
func NewPntContractRepository(db *gorm.DB) PntContractRepository {
	return &pntContractConnection{
		connection: db,
	}
}
