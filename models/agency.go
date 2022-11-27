package models

type Agency struct {
	ID                   uint
	AgencyLevelCode      string
	FullTimeLevelCode    string
	PntAgencyLevelCode   string
	PntFullTimeLevelCode string
}

func (Agency) TableName() string {
	return "agencies"
}
