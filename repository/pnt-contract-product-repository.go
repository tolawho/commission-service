package repository

import (
	"gorm.io/gorm"
	"medici.vn/commission-serivce/models"
)

// PntContractProductRepository  is contract what pntContractProductRepository can do to db
type PntContractProductRepository interface {
	FindByContractId(ContractId uint) []*models.PntContractProduct
}

type pntContractProductConnection struct {
	connection *gorm.DB
}

func (db pntContractProductConnection) FindByContractId(ContractId uint) []*models.PntContractProduct {
	var pntContractProducts []*models.PntContractProduct
	db.connection.Where("pnt_contract_id = ?", ContractId).Find(&pntContractProducts)
	return pntContractProducts
}

// NewPntContractProductRepository is creates a new instance of PntContractProductRepository
func NewPntContractProductRepository(db *gorm.DB) PntContractProductRepository {
	return &pntContractProductConnection{
		connection: db,
	}
}
