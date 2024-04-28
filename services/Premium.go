package services

import (
	"funeral_parlour/models"
	"funeral_parlour/models/dto"
	"funeral_parlour/repository/premium"
)

type Premium struct {
	svc premium.PremiumRepository
}

func NewPremiumService(svc premium.PremiumRepository) PremiumService {
	return &Premium{svc: svc}
}

func (p Premium) AddPremium(premium dto.FPPremium) (models.FPPremium, error) {
	return p.svc.AddPremium(models.FPPremium{
		FPPremiumName:        premium.FPPremiumName,
		FPPremiumDescription: premium.FPPremiumDescription,
		FPPremiumServiceID:   premium.FPPremiumServiceID,
		FPPremiumStatus:      premium.FPPremiumStatus,
	})
}

func (p Premium) UpdatePremium(premium dto.FPPremium) (models.FPPremium, error) {
	return p.svc.UpdatePremium(models.FPPremium{
		FPPremiumID:          premium.FPPremiumID,
		FPPremiumName:        premium.FPPremiumName,
		FPPremiumDescription: premium.FPPremiumDescription,
		FPPremiumServiceID:   premium.FPPremiumServiceID,
		FPPremiumStatus:      premium.FPPremiumStatus,
	})
}

func (p Premium) FindAllPremium() ([]models.FPPremium, error) {
	return p.svc.FindAllPremium()
}

func (p Premium) DeletePremium(premiumId int) error {
	return p.svc.DeletePremium(premiumId)
}

func (p Premium) FindPremium(premiumId int) (models.FPPremium, error) {
	return p.svc.FindPremium(premiumId)
}
