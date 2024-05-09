package controllers

import (
	"funeral_parlour/controllers/helper"
	"funeral_parlour/models"
	"funeral_parlour/models/dto"
	"funeral_parlour/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Branch struct {
	svc services.BranchService
}

func NewBranchController(svc services.BranchService) Branch {
	return Branch{svc: svc}
}

// AddBranch godoc
// @Summary      This API Creates a new branch
// @Description  This API Creates a new branch
// @Tags         Branch
// @Accept       json
// @Produce      json
// @Param premium body models.FPBranch true "New Branch"
// @Success      200  {object}  models.FPBranch
// @Failure      400  {object}  dto.FPAPIResponse
// @Failure      404  {object}  dto.FPAPIResponse
// @Failure      500  {object}  dto.FPAPIResponse
// @Router       /branch/add [POST]
func (p *Branch) AddBranch(c *gin.Context) {
	var branch models.FPBranch

	if err := c.ShouldBindJSON(&branch); err != nil {
		_err := helper.ValidateStruct(err)
		if _err != "" {
			c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: _err, StatusCode: http.StatusBadRequest})
			return
		}

	}

	resp, err := p.svc.AddBranch(branch)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	c.JSON(http.StatusOK, resp)

}

// UpdateBranch godoc
// @Summary      This API Update existing branch
// @Description  This API Update existing branch
// @Tags         Branch
// @Accept       json
// @Produce      json
// @Param branch body models.FPBranch true "Update Branch"
// @Success      200  {object}  models.FPBranch
// @Failure      400  {object}  dto.FPAPIResponse
// @Failure      404  {object}  dto.FPAPIResponse
// @Failure      500  {object}  dto.FPAPIResponse
// @Router       /branch/update [PUT]
func (p *Branch) UpdateBranch(c *gin.Context) {
	var branch models.FPBranch

	if err := c.ShouldBindJSON(&branch); err != nil {
		_err := helper.ValidateStruct(err)
		if _err != "" {
			c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: _err, StatusCode: http.StatusBadRequest})
			return
		}

	}

	resp, err := p.svc.UpdateBranch(branch)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// FindAllBranch godoc
// @Summary      This API List all Branches for a Service Provider
// @Description  This API List all Branches for a Service Provider
// @Tags         Branch
// @Accept       json
// @Produce      json
// @Param spId path string true "The id of Service Provider"
// @Success      200  {object}  []models.FPBranch
// @Failure      400  {object}  dto.FPAPIResponse
// @Failure      404  {object}  dto.FPAPIResponse
// @Failure      500  {object}  dto.FPAPIResponse
// @Router       /branch/all/{spId} [GET]
func (p *Branch) FindAllBranch(c *gin.Context) {

	spId := c.Param("spId")
	resp, err := p.svc.FindAllBranch(spId)

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, resp)

}

// DeleteBranch godoc
// @Summary      This API Delete branch by ID
// @Description  This API Delete branch by ID
// @Tags         Branch
// @Accept       json
// @Produce      json
// @Param branchId path string true "The id of Branch"
// @Param spId path string true "The id of ServiceProvider"
// @Success      200  {object}  dto.FPAPIResponse
// @Failure      400  {object}  dto.FPAPIResponse
// @Failure      404  {object}  dto.FPAPIResponse
// @Failure      500  {object}  dto.FPAPIResponse
// @Router       /branch/delete/{branchId}/{spId} [DELETE]
func (p *Branch) DeleteBranch(c *gin.Context) {
	branchId := c.Param("branchId")
	spId := c.Param("spId")

	err := p.svc.DeleteBranch(branchId, spId)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, dto.FPAPIResponse{Message: "Success", StatusCode: http.StatusOK})
}

// FindBranch godoc
// @Summary      This API Finds branch by ID
// @Description  This API Finds branch by ID
// @Tags         Branch
// @Accept       json
// @Produce      json
// @Param branchId path string true "The id of branch"
// @Param spId path string true "The id of ServiceProvider"
// @Success      200  {object}  models.FPBranch
// @Failure      400  {object}  dto.FPAPIResponse
// @Failure      404  {object}  dto.FPAPIResponse
// @Failure      500  {object}  dto.FPAPIResponse
// @Router       /branch/find/{branchId}/{spId} [GET]
func (p *Branch) FindBranch(c *gin.Context) {

	branchId := c.Param("branchId")
	spId := c.Param("spId")

	resp, err := p.svc.FindBranch(branchId, spId)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.FPAPIResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, resp)

}
