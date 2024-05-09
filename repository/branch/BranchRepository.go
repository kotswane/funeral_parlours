package branch

import "funeral_parlour/models"

type BranchRepository interface {
	AddBranch(branch models.FPBranch) (models.FPBranch, error)
	UpdateBranch(branch models.FPBranch) (models.FPBranch, error)
	FindAllBranch(spId string) ([]models.FPBranch, error)
	DeleteBranch(branchId string, spId string) error
	FindBranch(branchId string, spId string) (models.FPBranch, error)
}
