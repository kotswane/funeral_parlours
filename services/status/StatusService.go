package status

import (
	"funeral_parlour/models"
	"funeral_parlour/models/dto"
)

type StatusService interface {
	AddStatus(status dto.FPStatus) (models.FPStatus, error)
	UpdateStatus(status dto.FPStatus) (models.FPStatus, error)
	FindAllStatus() ([]models.FPStatus, error)
	DeleteStatus(status string) error
	FindStatus(status string) (models.FPStatus, error)
}
