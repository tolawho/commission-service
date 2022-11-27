package repository

import (
	"gorm.io/gorm"
	"medici.vn/commission-serivce/models"
)

// PntCommissionFormulaRepository  is contract what pntProductRepository can do to db
type PntCommissionFormulaRepository interface {
	FindFormula(condition models.PntCommissionFormula) *models.PntCommissionFormula
	FindFormulas(condition models.PntCommissionFormula) []*models.PntCommissionFormula
}

type pntCommissionFormulaConnection struct {
	connection *gorm.DB
}

func (db pntCommissionFormulaConnection) FindFormulas(condition models.PntCommissionFormula) []*models.PntCommissionFormula {
	var formulas []*models.PntCommissionFormula
	db.connection.Where(&condition).Find(&formulas)
	return formulas
}

func (db pntCommissionFormulaConnection) FindFormula(condition models.PntCommissionFormula) *models.PntCommissionFormula {
	var formula *models.PntCommissionFormula
	db.connection.Where(&condition).First(&formula)
	return formula
}

// NewPntCommissionFormulaRepository is creates a new instance of PntCommissionFormulaRepository
func NewPntCommissionFormulaRepository(db *gorm.DB) PntCommissionFormulaRepository {
	return &pntCommissionFormulaConnection{
		connection: db,
	}
}
