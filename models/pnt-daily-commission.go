package models

import (
	"gorm.io/gorm"
	"time"
)

type PntDailyCommission struct {
	gorm.Model
	Day                time.Time
	AgencyId           uint
	PntContractId      uint
	SourceId           uint
	ReferralCommission float32
	Amount             float32
	Note               string
	LevelCode          string
	SourceLevelCode    string
	SourceModel        string
	SysCommission      float32
	SysRate            float32
	SysTaxRate         float32
	IsOldData          bool
	DatePayment        time.Time
	PolicyId           uint
	PntContract        PntContract `gorm:"foreignKey:PntContractId" json:"pntContract"`
}

func (PntDailyCommission) TableName() string {
	return "pnt_daily_commissions"
}
