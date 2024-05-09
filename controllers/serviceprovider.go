package controllers

import (
	"funeral_parlour/controllers/helper"
	"funeral_parlour/models/dto"
	"funeral_parlour/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceProvider struct {
	svc services.ServiceProviderService
}

func NewServiceProviderController(svc services.ServiceProviderService) ServiceProvider {
	return ServiceProvider{svc: svc}
}

// AddServiceProvider godoc
// @Summary      This API Creates a new Service Provider
// @Description  This API Creates a new Service Provider
// @Tags         Service Provider
// @Accept       json
// @Produce      json
// @Param serviceprovider body models.FPServiceProvider true "New Service Provider"
// @Success      200  {object}  models.FPServiceProvider
// @Failure      400  {object}  dto.FPAPIResponse
// @Failure      404  {object}  dto.FPAPIResponse
// @Failure      500  {object}  dto.FPAPIResponse
// @Router       /serviceprovider/add [POST]
func (s *ServiceProvider) AddServiceProvider(c *gin.Context) {
	var serviceprovider dto.FPServiceProvider

	if err := c.ShouldBindJSON(&serviceprovider); err != nil {
		_err := helper.ValidateStruct(err)
		if _err != "" {
			c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: _err, StatusCode: http.StatusBadRequest})
			return
		}
	}

	resp, err := s.svc.AddServiceProvider(serviceprovider)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateServiceProvider godoc
// @Summary      This API Updates a Service Provider
// @Description  This API Updates a Service Provider
// @Tags         Service Provider
// @Accept       json
// @Produce      json
// @Param serviceprovider body models.FPServiceProvider true "Existing Service Provider"
// @Success      200  {object}  models.FPServiceProvider
// @Failure      400  {object}  dto.FPAPIResponse
// @Failure      404  {object}  dto.FPAPIResponse
// @Failure      500  {object}  dto.FPAPIResponse
// @Router       /serviceprovider/update [POST]
func (s *ServiceProvider) UpdateServiceProvider(c *gin.Context) {
	var serviceprovider dto.FPServiceProvider

	if err := c.ShouldBindJSON(&serviceprovider); err != nil {
		_err := helper.ValidateStruct(err)
		if _err != "" {
			c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: _err, StatusCode: http.StatusBadRequest})
			return
		}
	}

	resp, err := s.svc.UpdateServiceProvider(serviceprovider)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// FindAllServiceProvider godoc
// @Summary      This API List all Service Providers
// @Description  This API List all Service Providers
// @Tags         Service Provider
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.FPServiceProvider
// @Failure      400  {object}  dto.FPAPIResponse
// @Failure      404  {object}  dto.FPAPIResponse
// @Failure      500  {object}  dto.FPAPIResponse
// @Router       /serviceprovider [GET]
func (s *ServiceProvider) FindAllServiceProvider(c *gin.Context) {

	resp, err := s.svc.FindAllServiceProvider()

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteServiceProvider godoc
// @Summary      This API Delete Service Provider by ID
// @Description  This API Delete Service Provider by ID
// @Tags         Service Provider
// @Accept       json
// @Produce      json
// @Param serviceProviderId path string true "The id of ServiceProvider"
// @Success      200  {object}  dto.FPAPIResponse
// @Failure      400  {object}  dto.FPAPIResponse
// @Failure      404  {object}  dto.FPAPIResponse
// @Failure      500  {object}  dto.FPAPIResponse
// @Router       /serviceprovider/delete/{serviceproviderId} [DELETE]
func (s *ServiceProvider) DeleteServiceProvider(c *gin.Context) {
	serviceProviderId := c.Param("serviceProviderId")
	err := s.svc.DeleteServiceProvider(serviceProviderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, dto.FPAPIResponse{Message: "Success", StatusCode: http.StatusOK})
}

// FindServiceProvider godoc
// @Summary      This API Finds Service Provider by ID
// @Description  This API Finds Service Provider by ID
// @Tags         Service Provider
// @Accept       json
// @Produce      json
// @Param serviceProviderId path string true "The id of Service Provider"
// @Success      200  {object}  models.FPPremium
// @Failure      400  {object}  dto.FPAPIResponse
// @Failure      404  {object}  dto.FPAPIResponse
// @Failure      500  {object}  dto.FPAPIResponse
// @Router       /serviceprovider/{serviceproviderId} [GET]
func (s *ServiceProvider) FindServiceProvider(c *gin.Context) {
	serviceProviderId := c.Param("serviceProviderId")
	serviceProviderList, err := s.svc.FindServiceProvider(serviceProviderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, serviceProviderList)
}
