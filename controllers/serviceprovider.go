package controllers

import (
	"funeral_parlour/repository/serviceprovider"
	"github.com/gin-gonic/gin"
)

type ServiceProvider struct {
	svc serviceprovider.ServiceProvider
}

func NewServiceProviderController(svc serviceprovider.ServiceProvider) ServiceProvider {
	return ServiceProvider{svc: svc}
}

func (s *ServiceProvider) AddServiceProvider(c *gin.Context) {

}

func (s *ServiceProvider) UpdateServiceProvider(c *gin.Context) {

}

func (s *ServiceProvider) FindAllServiceProvider(c *gin.Context) {

}

func (s *ServiceProvider) DeleteServiceProvider(c *gin.Context) {

}

func (s *ServiceProvider) FindServiceProvider(c *gin.Context) {

}
