package repository

import (
	"gorm.io/gorm"
	pntPolicy "medici.vn/commission-serivce/enums/pnt-policy"
	"medici.vn/commission-serivce/models"
)

// PntPolicyRepository  is contract what pntPolicyRepository can do to db
type PntPolicyRepository interface {
	FindActive() models.PntPolicy
}

type pntPolicyConnection struct {
	connection *gorm.DB
}

func (db pntPolicyConnection) FindActive() models.PntPolicy {
	var policy = models.PntPolicy{}
	db.connection.Where("status = ?", pntPolicy.ON).Take(&policy)
	return policy
}

// NewPntPolicyRepository is creates a new instance of PntPolicyRepository
func NewPntPolicyRepository(db *gorm.DB) PntPolicyRepository {
	return &pntPolicyConnection{
		connection: db,
	}
}
