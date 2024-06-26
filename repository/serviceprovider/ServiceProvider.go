package serviceprovider

import (
	"errors"
	"funeral_parlour/models"

	"gorm.io/gorm"
)

type ServiceProvider struct {
	db *gorm.DB
}

func NewServiceProvider(db *gorm.DB) ServiceProviderRepository {
	return &ServiceProvider{db: db}
}

func (s ServiceProvider) AddServiceProvider(serviceProvider models.FPServiceProvider) (models.FPServiceProvider, error) {

	if _, err := s.FindServiceProviderFSP(serviceProvider.FPServiceProviderFspNumber); err == nil {
		response := s.db.Create(serviceProvider)
		if response.Error != nil {
			return models.FPServiceProvider{}, response.Error
		}
		return serviceProvider, nil
	}

	return serviceProvider, errors.New("service provider already exist")
}

func (s ServiceProvider) UpdateServiceProvider(serviceProvider models.FPServiceProvider) (models.FPServiceProvider, error) {
	serviceProviderUpdate := models.FPServiceProvider{
		FPServiceProviderName:          serviceProvider.FPServiceProviderName,
		FPServiceProviderFspNumber:     serviceProvider.FPServiceProviderFspNumber,
		FPServiceProviderAddress:       serviceProvider.FPServiceProviderAddress,
		FPServiceProviderEmail:         serviceProvider.FPServiceProviderEmail,
		FPServiceProviderContactNumber: serviceProvider.FPServiceProviderContactNumber,
		FPServiceProviderStatus:        serviceProvider.FPServiceProviderStatus,
	}

	if err := s.db.Model(&models.FPServiceProvider{}).Where("sp_id = ?", serviceProvider.FPServiceProviderID).Updates(serviceProviderUpdate).Error; err != nil {
		return models.FPServiceProvider{}, err
	}
	return serviceProviderUpdate, nil
}

func (s ServiceProvider) FindAllServiceProvider() ([]models.FPServiceProvider, error) {
	var serviceProviderList []models.FPServiceProvider
	if err := s.db.Find(&serviceProviderList).Error; err != nil {
		return serviceProviderList, err
	}
	return serviceProviderList, nil
}

func (s ServiceProvider) DeleteServiceProvider(serviceProviderId string) error {
	if err := s.db.Where("sp_id = ?", serviceProviderId).Delete(&models.FPServiceProvider{}).Error; err != nil {
		return err
	}
	return nil
}

func (s ServiceProvider) FindServiceProvider(serviceProviderId string) (models.FPServiceProvider, error) {
	var serviceProvider models.FPServiceProvider
	if err := s.db.Where("sp_id = ?", serviceProviderId).First(&serviceProvider).Error; err != nil {
		return serviceProvider, err
	}
	return serviceProvider, nil
}

func (s ServiceProvider) FindServiceProviderFSP(fsp string) (models.FPServiceProvider, error) {
	var serviceProvider models.FPServiceProvider
	if err := s.db.Where("sp_fsp_number = ?", fsp).First(&serviceProvider).Error; err != nil {
		return serviceProvider, err
	}
	return serviceProvider, nil
}
