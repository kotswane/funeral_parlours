package models

type FPStatus struct {
	FPStatusID   uint   `gorm:"column:fp_status_id;primaryKey;autoIncrement"`
	FPStatusName string `gorm:"column:fp_status_name;not null;default:''"`
}

func (m FPStatus) TableName() string {
	return "fp_status"
}
