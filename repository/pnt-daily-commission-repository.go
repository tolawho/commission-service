package repository

import (
	"gorm.io/gorm"
	"medici.vn/commission-serivce/models"
)

// PntDailyCommissionRepository  is contract what pntDailyCommissionRepository can do to db
type PntDailyCommissionRepository interface {
	FirstOrCreate(condition models.PntDailyCommission, commission models.PntDailyCommission)
}

type pntDailyCommissionConnection struct {
	connection *gorm.DB
}

func (db pntDailyCommissionConnection) FirstOrCreate(
	condition models.PntDailyCommission,
	commission models.PntDailyCommission) {
	db.connection.Where(&condition).FirstOrCreate(&commission)
}

// NewPntDailyCommissionRepository is creates a new instance of PntDailyCommissionRepository
func NewPntDailyCommissionRepository(db *gorm.DB) PntDailyCommissionRepository {
	return &pntDailyCommissionConnection{
		connection: db,
	}
}
