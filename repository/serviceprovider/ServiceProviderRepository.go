package serviceprovider

import "funeral_parlour/models"

type ServiceProviderRepository interface {
	AddServiceProvider(serviceProvider models.FPServiceProvider) (models.FPServiceProvider, error)
	UpdateServiceProvider(serviceProvider models.FPServiceProvider) (models.FPServiceProvider, error)
	FindAllServiceProvider() ([]models.FPServiceProvider, error)
	DeleteServiceProvider(serviceProviderId string) error
	FindServiceProvider(serviceProviderId string) (models.FPServiceProvider, error)

	FindServiceProviderFSP(fsp string) (models.FPServiceProvider, error)
}
