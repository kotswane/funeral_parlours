package services

import (
	"funeral_parlour/models"
	"funeral_parlour/models/dto"
)

type PremiumService interface {
	AddPremium(premium dto.FPCreatePremium) (models.FPPremium, error)
	UpdatePremium(premium dto.FPUpdatePremium) (models.FPPremium, error)
	FindAllPremium() ([]models.FPPremium, error)
	DeletePremium(premiumId string, spId string) error
	FindPremium(premiumId string, spId string) (models.FPPremium, error)
	FindPremiumBySP(spId string) ([]models.FPPremium, error)
}
