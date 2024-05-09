package controllers

import (
	"funeral_parlour/controllers/helper"
	"funeral_parlour/models/dto"
	"funeral_parlour/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Premium struct {
	svc services.PremiumService
}

func NewPremiumController(svc services.PremiumService) Premium {
	return Premium{svc: svc}
}

// AddPremium godoc
// @Summary      This API Creates a new premium
// @Description  This API Creates a new premium
// @Tags         Premium
// @Accept       json
// @Produce      json
// @Param premium body dto.FPCreatePremium true "New Premium"
// @Success      200  {object}  models.FPPremium
// @Failure      400  {object}  dto.FPAPIResponse
// @Failure      404  {object}  dto.FPAPIResponse
// @Failure      500  {object}  dto.FPAPIResponse
// @Router       /premium/add [POST]
func (p *Premium) AddPremium(c *gin.Context) {
	var premium dto.FPCreatePremium

	if err := c.ShouldBindJSON(&premium); err != nil {
		_err := helper.ValidateStruct(err)
		if _err != "" {
			c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: _err, StatusCode: http.StatusBadRequest})
			return
		}

	}

	resp, err := p.svc.AddPremium(premium)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	c.JSON(http.StatusOK, resp)

}

// UpdatePremium godoc
// @Summary      This API Update existing premium
// @Description  This API Update existing premium
// @Tags         Premium
// @Accept       json
// @Produce      json
// @Param premium body dto.FPUpdatePremium true "Update Premium"
// @Success      200  {object}  models.FPPremium
// @Failure      400  {object}  dto.FPAPIResponse
// @Failure      404  {object}  dto.FPAPIResponse
// @Failure      500  {object}  dto.FPAPIResponse
// @Router       /premium/update [PUT]
func (p *Premium) UpdatePremium(c *gin.Context) {
	var premium dto.FPUpdatePremium

	if err := c.ShouldBindJSON(&premium); err != nil {
		_err := helper.ValidateStruct(err)
		if _err != "" {
			c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: _err, StatusCode: http.StatusBadRequest})
			return
		}

	}

	resp, err := p.svc.UpdatePremium(premium)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// FindAllPremium godoc
// @Summary      This API List all premiums
// @Description  This API List all premiums
// @Tags         Premium
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.FPPremium
// @Failure      400  {object}  dto.FPAPIResponse
// @Failure      404  {object}  dto.FPAPIResponse
// @Failure      500  {object}  dto.FPAPIResponse
// @Router       /premium [GET]
func (p *Premium) FindAllPremium(c *gin.Context) {

	resp, err := p.svc.FindAllPremium()

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, resp)

}

// DeletePremium godoc
// @Summary      This API Delete premium by ID
// @Description  This API Delete premium by ID
// @Tags         Premium
// @Accept       json
// @Produce      json
// @Param id path string true "The id of FPPremium"
// @Param spId path string true "The id of ServiceProvider"
// @Success      200  {object}  dto.FPAPIResponse
// @Failure      400  {object}  dto.FPAPIResponse
// @Failure      404  {object}  dto.FPAPIResponse
// @Failure      500  {object}  dto.FPAPIResponse
// @Router       /premium/delete/{id}/{spId} [DELETE]
func (p *Premium) DeletePremium(c *gin.Context) {
	id := c.Param("id")
	spId := c.Param("spId")

	err := p.svc.DeletePremium(id, spId)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, dto.FPAPIResponse{Message: "Success", StatusCode: http.StatusOK})
}

// FindPremium godoc
// @Summary      This API Finds premium by ID
// @Description  This API Finds premium by ID
// @Tags         Premium
// @Accept       json
// @Produce      json
// @Param id path string true "The id of FPPremium"
// @Param spId path string true "The id of FPPremiumServiceProvider"
// @Success      200  {object}  models.FPPremium
// @Failure      400  {object}  dto.FPAPIResponse
// @Failure      404  {object}  dto.FPAPIResponse
// @Failure      500  {object}  dto.FPAPIResponse
// @Router       /premium/find/{id}/{spId} [GET]
func (p *Premium) FindPremium(c *gin.Context) {

	id := c.Param("id")
	spId := c.Param("spId")

	resp, err := p.svc.FindPremium(id, spId)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, resp)

}

// FindPremiumBySP godoc
// @Summary      This API Finds premium by ServiceProviderID
// @Description  This API Finds premium by ServiceProviderID
// @Tags         Premium
// @Accept       json
// @Produce      json
// @Param spId path string true "The id of FPPremiumServiceProvider"
// @Success      200  {object}  []models.FPPremium
// @Failure      400  {object}  dto.FPAPIResponse
// @Failure      404  {object}  dto.FPAPIResponse
// @Failure      500  {object}  dto.FPAPIResponse
// @Router       /premium/findbysp/{spId} [GET]
func (p *Premium) FindPremiumBySP(c *gin.Context) {

	spId := c.Param("spId")

	resp, err := p.svc.FindPremiumBySP(spId)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, resp)

}
