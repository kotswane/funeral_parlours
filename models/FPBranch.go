package models

type FPBranch struct {
	FPBranchID                string `gorm:"column:fp_sp_branch_id;primaryKey"`
	FPBranchServiceProviderId string `gorm:"column:fp_sp_id"`
	FPBranchName              string `gorm:"column:fp_sp_branch"`
}
