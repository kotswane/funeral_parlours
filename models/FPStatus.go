package models

type FPStatus struct {
	FPStatusID   string `gorm:"column:fp_status_id;primaryKey"`
	FPStatusName string `gorm:"column:fp_status_name;not null;default:''"`
}

func (m FPStatus) TableName() string {
	return "fp_status"
}
