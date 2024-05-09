package models

type FPMemberType struct {
	FPMemberTypeID          string `gorm:"column:fp_member_type_id;primaryKey"`
	FPMemberTypeDescription string `gorm:"column:fp_member_type_description"`
}

func (m FPMemberType) TableName() string {
	return "fp_member_types"
}
