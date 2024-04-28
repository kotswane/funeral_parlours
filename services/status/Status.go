package status

import (
	"funeral_parlour/models"
	"funeral_parlour/models/dto"
	"funeral_parlour/repository/status"
)

type Status struct {
	svc status.StatusRepository
}

func NewStatus(svc status.StatusRepository) StatusService {
	return &Status{svc: svc}
}

func (s Status) AddStatus(status dto.FPStatus) (models.FPStatus, error) {
	return s.svc.AddStatus(models.FPStatus{
		FPStatusName: status.FPStatusName,
	})
}

func (s Status) UpdateStatus(status dto.FPStatus) (models.FPStatus, error) {
	return s.svc.UpdateStatus(models.FPStatus{
		FPStatusID:   status.FPStatusID,
		FPStatusName: status.FPStatusName,
	})
}

func (s Status) FindAllStatus() ([]models.FPStatus, error) {
	return s.svc.FindAllStatus()
}

func (s Status) DeleteStatus(statusId int) error {
	return s.svc.DeleteStatus(statusId)
}

func (s Status) FindStatus(statusId int) (models.FPStatus, error) {
	return s.svc.FindStatus(statusId)
}
