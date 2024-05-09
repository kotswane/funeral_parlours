package services

import (
	"funeral_parlour/models"
	"funeral_parlour/models/dto"
)

type ServiceProviderService interface {
	AddServiceProvider(serviceProvider dto.FPServiceProvider) (models.FPServiceProvider, error)
	UpdateServiceProvider(serviceProvider dto.FPServiceProvider) (models.FPServiceProvider, error)
	FindAllServiceProvider() ([]models.FPServiceProvider, error)
	DeleteServiceProvider(serviceProviderId string) error
	FindServiceProvider(serviceProviderId string) (models.FPServiceProvider, error)
}
