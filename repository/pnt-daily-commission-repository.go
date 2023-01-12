package repository

import (
	"gorm.io/gorm"
	"medici.vn/commission-serivce/models"
)

// PntDailyCommissionRepository  is contract what pntDailyCommissionRepository can do to db
type PntDailyCommissionRepository interface {
	FirstOrCreate(condition models.PntDailyCommission, commission models.PntDailyCommission) (any, error)
	Update(condition models.PntDailyCommission, commission models.PntDailyCommission) (models.PntDailyCommission, error)
}

type pntDailyCommissionConnection struct {
	connection *gorm.DB
}

func (db pntDailyCommissionConnection) Update(
	condition models.PntDailyCommission,
	commission models.PntDailyCommission) (models.PntDailyCommission, error) {
	result := db.connection.Where(&condition).Updates(&commission)
	return commission, result.Error
}

func (db pntDailyCommissionConnection) FirstOrCreate(
	condition models.PntDailyCommission,
	commission models.PntDailyCommission) (any, error) {
	result := db.connection.Where(&condition).FirstOrCreate(&commission)
	return commission, result.Error
}

// NewPntDailyCommissionRepository is creates a new instance of PntDailyCommissionRepository
func NewPntDailyCommissionRepository(db *gorm.DB) PntDailyCommissionRepository {
	return &pntDailyCommissionConnection{
		connection: db,
	}
}
