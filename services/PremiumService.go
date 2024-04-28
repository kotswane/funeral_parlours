package services

import (
	"funeral_parlour/models"
	"funeral_parlour/models/dto"
)

type PremiumService interface {
	AddPremium(premium dto.FPPremium) (models.FPPremium, error)
	UpdatePremium(premium dto.FPPremium) (models.FPPremium, error)
	FindAllPremium() ([]models.FPPremium, error)
	DeletePremium(premiumId int) error
	FindPremium(premiumId int) (models.FPPremium, error)
}
