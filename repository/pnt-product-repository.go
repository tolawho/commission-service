package repository

import (
	"gorm.io/gorm"
)

// PntProductRepository  is contract what pntProductRepository can do to db
type PntProductRepository interface {
}

type pntProductConnection struct {
	connection *gorm.DB
}

// NewPntProductRepository is creates a new instance of PntProductRepository
func NewPntProductRepository(db *gorm.DB) PntProductRepository {
	return &pntProductConnection{
		connection: db,
	}
}
