package models

import (
	"gorm.io/gorm"
	pnt_transaction "medici.vn/commission-serivce/enums/pnt-transaction"
)

type PntTransaction struct {
	gorm.Model
	Note          string
	AgencyId      uint
	PntContractId uint
	Type          pnt_transaction.PntTransactionType
	Status        pnt_transaction.PntTransactionStatus
	Amount        float32
}

func (PntTransaction) TableName() string {
	return "pnt_transactions"
}
