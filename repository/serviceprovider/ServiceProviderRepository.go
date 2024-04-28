package serviceprovider

import "funeral_parlour/models"

type ServiceProviderRepository interface {
	AddServiceProvider(serviceProvider models.FPServiceProvider) (models.FPServiceProvider, error)
	UpdateServiceProvider(serviceProvider models.FPServiceProvider) (models.FPServiceProvider, error)
	FindAllServiceProvider() ([]models.FPServiceProvider, error)
	DeleteServiceProvider(serviceProviderId int) error
	FindServiceProvider(serviceProviderId int) (models.FPServiceProvider, error)
}
