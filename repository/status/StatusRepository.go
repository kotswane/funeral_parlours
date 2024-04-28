package status

import "funeral_parlour/models"

type StatusRepository interface {
	AddStatus(status models.FPStatus) (models.FPStatus, error)
	UpdateStatus(status models.FPStatus) (models.FPStatus, error)
	FindAllStatus() ([]models.FPStatus, error)
	DeleteStatus(statusId int) error
	FindStatus(statusId int) (models.FPStatus, error)
}
