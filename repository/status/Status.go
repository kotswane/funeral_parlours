package status

import (
	"funeral_parlour/models"

	"gorm.io/gorm"
)

type Status struct {
	db *gorm.DB
}

func NewStatus(db *gorm.DB) StatusRepository {
	return &Status{db: db}
}

func (s Status) AddStatus(status models.FPStatus) (models.FPStatus, error) {
	results := s.db.Create(status)
	if results.Error != nil {
		return models.FPStatus{}, results.Error
	}
	return status, nil
}

func (s Status) UpdateStatus(status models.FPStatus) (models.FPStatus, error) {
	if err := s.db.Model(&models.FPStatus{}).Where("fp_status_id = ?", status.FPStatusID).Update("fp_status_name", status.FPStatusName).Error; err != nil {
		return models.FPStatus{}, err
	}
	return status, nil
}

func (s Status) FindAllStatus() ([]models.FPStatus, error) {
	var statusList []models.FPStatus
	if err := s.db.Find(&statusList).Error; err != nil {
		return statusList, err
	}
	return statusList, nil
}

func (s Status) DeleteStatus(statusId string) error {
	if err := s.db.Where("fp_status_id = ?", statusId).Delete(&models.FPStatus{}).Error; err != nil {
		return err
	}
	return nil
}

func (s Status) FindStatus(statusId string) (models.FPStatus, error) {
	var status models.FPStatus
	if err := s.db.Where("fp_status_id = ?", statusId).First(&status).Error; err != nil {
		return status, err
	}
	return status, nil
}
