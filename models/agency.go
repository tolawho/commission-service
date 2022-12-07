package models

type Agency struct {
	ID                uint
	Name              string
	Code              string
	LvPartTime        string
	LvPartTimePlus    string
	LvFullTime        string
	PntLvPartTime     string
	PntLvPartTimePlus string
	PntLvFullTime     string
}

func (Agency) TableName() string {
	return "agencies"
}
