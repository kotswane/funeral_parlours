package services

import (
	"errors"
	"funeral_parlour/models"
	"funeral_parlour/repository/branch"
	"funeral_parlour/services/helper"

	"github.com/google/uuid"
)

type Branch struct {
	repo branch.BranchRepository
}

func NewBranchService(repo branch.BranchRepository) BranchService {
	return &Branch{repo: repo}
}

// AddBranch implements BranchService.
func (b *Branch) AddBranch(branch models.FPBranch) (models.FPBranch, error) {
	isValidGuid := helper.IsValidUUID(branch.FPBranchServiceProviderId)
	if !isValidGuid {
		return models.FPBranch{}, errors.New("invalid FPBranchServiceProviderId")
	}

	return b.repo.AddBranch(models.FPBranch{
		FPBranchID:                uuid.NewString(),
		FPBranchServiceProviderId: branch.FPBranchServiceProviderId,
		FPBranchName:              branch.FPBranchName,
	})
}

// DeleteBranch implements BranchService.
func (b *Branch) DeleteBranch(branchId string, spId string) error {
	isValidGuid := helper.IsValidUUID(branchId)
	if !isValidGuid {
		return errors.New("invalid branchId")
	}

	isValidGuid = helper.IsValidUUID(spId)
	if !isValidGuid {
		return errors.New("invalid serviceProviderId")
	}

	return b.repo.DeleteBranch(branchId, spId)
}

// FindAllBranch implements BranchService.
func (b *Branch) FindAllBranch(spId string) ([]models.FPBranch, error) {

	isValidGuid := helper.IsValidUUID(spId)
	if !isValidGuid {
		return []models.FPBranch{}, errors.New("invalid serviceProviderId")
	}
	return b.repo.FindAllBranch(spId)
}

// FindBranch implements BranchService.
func (b *Branch) FindBranch(branchId string, spId string) (models.FPBranch, error) {
	isValidGuid := helper.IsValidUUID(branchId)
	if !isValidGuid {
		return models.FPBranch{}, errors.New("invalid branchId")
	}

	isValidGuid = helper.IsValidUUID(spId)
	if !isValidGuid {
		return models.FPBranch{}, errors.New("invalid serviceProviderId")
	}
	return b.repo.FindBranch(branchId, spId)
}

// UpdateBranch implements BranchService.
func (b *Branch) UpdateBranch(branch models.FPBranch) (models.FPBranch, error) {
	isValidGuid := helper.IsValidUUID(branch.FPBranchID)
	if !isValidGuid {
		return models.FPBranch{}, errors.New("invalid branchId")
	}
	return b.repo.UpdateBranch(branch)
}
