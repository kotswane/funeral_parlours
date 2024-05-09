package status

import (
	"errors"
	"funeral_parlour/models"
	"funeral_parlour/models/dto"
	"funeral_parlour/repository/status"
	"funeral_parlour/services/helper"
)

type Status struct {
	repo status.StatusRepository
}

func NewStatus(repo status.StatusRepository) StatusService {
	return &Status{repo: repo}
}

func (s Status) AddStatus(status dto.FPStatus) (models.FPStatus, error) {
	return s.repo.AddStatus(models.FPStatus{
		FPStatusName: status.FPStatusName,
	})
}

func (s Status) UpdateStatus(status dto.FPStatus) (models.FPStatus, error) {
	isValidGuid := helper.IsValidUUID(status.FPStatusID)
	if !isValidGuid {
		return models.FPStatus{}, errors.New("invalid StatusID")
	}

	return s.repo.UpdateStatus(models.FPStatus{
		FPStatusID:   status.FPStatusID,
		FPStatusName: status.FPStatusName,
	})
}

func (s Status) FindAllStatus() ([]models.FPStatus, error) {
	return s.repo.FindAllStatus()
}

func (s Status) DeleteStatus(statusId string) error {
	isValidGuid := helper.IsValidUUID(statusId)
	if !isValidGuid {
		return errors.New("invalid StatusID")
	}
	return s.repo.DeleteStatus(statusId)
}

func (s Status) FindStatus(statusId string) (models.FPStatus, error) {
	isValidGuid := helper.IsValidUUID(statusId)
	if !isValidGuid {
		return models.FPStatus{}, errors.New("invalid StatusID")
	}
	return s.repo.FindStatus(statusId)
}
