package repository

import (
	"gorm.io/gorm"
	"medici.vn/commission-serivce/models"
)

// PntTransactionHistoryRepository is contract what pntTransactionHistoryRepository can do to db
type PntTransactionHistoryRepository interface {
	Update(condition models.PntTransactionHistory, transaction models.PntTransactionHistory) (models.PntTransactionHistory, error)
	Create(transaction models.PntTransactionHistory) (models.PntTransactionHistory, error)
}

type pntTransactionHistoryConnection struct {
	connection *gorm.DB
}

func (db pntTransactionHistoryConnection) Create(transaction models.PntTransactionHistory) (models.PntTransactionHistory, error) {
	result := db.connection.Create(&transaction)
	return transaction, result.Error
}

func (db pntTransactionHistoryConnection) Update(condition models.PntTransactionHistory, transaction models.PntTransactionHistory) (models.PntTransactionHistory, error) {
	result := db.connection.Where(&condition).Updates(&transaction)
	return transaction, result.Error
}

// NewPntTransactionHistoryRepository is creates a new instance of PntTransactionHistoryRepository
func NewPntTransactionHistoryRepository(db *gorm.DB) PntTransactionHistoryRepository {
	return &pntTransactionHistoryConnection{
		connection: db,
	}
}
