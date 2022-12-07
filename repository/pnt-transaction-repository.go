package repository

import (
	"gorm.io/gorm"
	"medici.vn/commission-serivce/models"
)

// PntTransactionRepository is contract what pntTransactionRepository can do to db
type PntTransactionRepository interface {
	FirstOrCreate(condition models.PntTransaction, transaction models.PntTransaction) (models.PntTransaction, error)
}

type pntTransactionConnection struct {
	connection *gorm.DB
}

func (db pntTransactionConnection) FirstOrCreate(condition models.PntTransaction, transaction models.PntTransaction) (models.PntTransaction, error) {
	result := db.connection.Where(&condition).FirstOrCreate(&transaction)
	return transaction, result.Error
}

// NewPntTransactionRepository is creates a new instance of PntTransactionRepository
func NewPntTransactionRepository(db *gorm.DB) PntTransactionRepository {
	return &pntTransactionConnection{
		connection: db,
	}
}
