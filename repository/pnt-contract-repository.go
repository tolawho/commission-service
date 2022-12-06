package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"medici.vn/commission-serivce/models"
)

// PntContractRepository  is contract what pntContractRepository can do to db
type PntContractRepository interface {
	First(pntContract models.PntContract) (models.PntContract, error)
}

type pntContractConnection struct {
	connection *gorm.DB
}

func (db pntContractConnection) First(pntContract models.PntContract) (models.PntContract, error) {
	if err := db.connection.Preload(clause.Associations).Where(&pntContract).First(&pntContract).Error; err != nil {
		return models.PntContract{}, err
	}
	return pntContract, nil
}

// NewPntContractRepository is creates a new instance of PntContractRepository
func NewPntContractRepository(db *gorm.DB) PntContractRepository {
	return &pntContractConnection{
		connection: db,
	}
}
