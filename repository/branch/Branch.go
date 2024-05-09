package branch

import (
	"funeral_parlour/models"

	"gorm.io/gorm"
)

type Branch struct {
	db *gorm.DB
}

func NewBranch(db *gorm.DB) BranchRepository {
	return &Branch{db: db}
}

// AddBranch implements BranchRepository.
func (b *Branch) AddBranch(branch models.FPBranch) (models.FPBranch, error) {
	response := b.db.Create(branch)
	if response.Error != nil {
		return models.FPBranch{}, response.Error
	}
	return branch, nil
}

// DeleteBranch implements BranchRepository.
func (b *Branch) DeleteBranch(branchId string, spId string) error {
	if err := b.db.Where("fp_sp_id = ? AND fp_sp_branch_id = ? ", spId, branchId).Delete(&models.FPBranch{}).Error; err != nil {
		return err
	}
	return nil
}

// FindAllBranch implements BranchRepository.
func (b *Branch) FindAllBranch(spId string) ([]models.FPBranch, error) {
	var branches []models.FPBranch
	if err := b.db.Where("fp_sp_id = ?", spId).Find(&branches).Error; err != nil {
		return branches, err
	}
	return branches, nil
}

// FindBranch implements BranchRepository.
func (b *Branch) FindBranch(branchId string, spId string) (models.FPBranch, error) {
	var branch models.FPBranch
	if err := b.db.Where("fp_sp_id = ? AND fp_sp_branch_id = ? ", spId, branchId).First(&branch).Error; err != nil {
		return branch, err
	}
	return branch, nil
}

// UpdateBranch implements BranchRepository.
func (b *Branch) UpdateBranch(branch models.FPBranch) (models.FPBranch, error) {
	updateBranch := models.FPBranch{
		FPBranchName: branch.FPBranchName,
	}

	if err := b.db.Model(&models.FPBranch{}).Where("fp_sp_branch_id = ? AND fp_sp_branch_id", branch.FPBranchID, branch.FPBranchServiceProviderId).Updates(&updateBranch).Error; err != nil {
		return models.FPBranch{}, err
	}
	return branch, nil
}
