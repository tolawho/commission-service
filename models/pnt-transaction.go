package models

import (
	"fmt"
	"gorm.io/gorm"
	pnt_transaction "medici.vn/commission-serivce/enums/pnt-transaction"
	"time"
)

type PntTransaction struct {
	gorm.Model
	Code          *string
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

func (p PntTransaction) AfterCreate(db *gorm.DB) error {
	db.Model(&p).Update("Code", fmt.Sprintf("GDPNT%s%d", time.Now().Format("20060102"), p.ID))
	return nil
}
