package dto

type FPMemberTypeCreate struct {
	FPMemberTypeID          string `json:"fp_member_type_id"`
	FPMemberTypeDescription string `gorm:"fp_member_type_description" validate:"required"`
}
type FPMemberTypeUpdate struct {
	FPMemberTypeID          string `json:"fp_member_type_id" validate:"required"`
	FPMemberTypeDescription string `gorm:"fp_member_type_description" validate:"required"`
}
