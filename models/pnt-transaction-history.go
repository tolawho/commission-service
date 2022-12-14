package models

import (
	"gorm.io/gorm"
	pnt_transaction "medici.vn/commission-serivce/enums/pnt-transaction"
)

type PntTransactionHistory struct {
	gorm.Model
	PntTransactionId uint
	PntContractId    uint
	AgencyId         uint
	Type             pnt_transaction.PntTransactionType
	Status           pnt_transaction.PntTransactionStatus
	Amount           float32
	Fee              float32
	Vat              float32
	Pit              float32
	Note             string
}

func (PntTransactionHistory) TableName() string {
	return "pnt_transaction_histories"
}
