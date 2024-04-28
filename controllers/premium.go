package controllers

import (
	"funeral_parlour/services"

	"github.com/gin-gonic/gin"
)

type Premium struct {
	svc services.PremiumService
}

func NewPremiumController(svc services.PremiumService) Premium {
	return Premium{svc: svc}
}

func (p *Premium) AddPremium(c *gin.Context) {

}

func (p *Premium) UpdatePremium(c *gin.Context) {

}

func (p *Premium) FindAllPremium(c *gin.Context) {

}

func (p *Premium) DeletePremium(c *gin.Context) {

}

func (p *Premium) FindPremium(c *gin.Context) {

}
