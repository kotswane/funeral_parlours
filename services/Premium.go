package services

import (
	"errors"
	"fmt"
	"funeral_parlour/models"
	"funeral_parlour/models/dto"
	"funeral_parlour/repository/premium"
	"funeral_parlour/services/helper"

	"github.com/google/uuid"
)

type Premium struct {
	repo premium.PremiumRepository
}

func NewPremiumService(repo premium.PremiumRepository) PremiumService {
	return &Premium{repo: repo}
}

func (p Premium) AddPremium(premium dto.FPCreatePremium) (models.FPPremium, error) {
	isValidGuid := helper.IsValidUUID(premium.FPPremiumServiceID)
	if !isValidGuid {
		return models.FPPremium{}, errors.New("invalid FPPremiumServiceID")
	}
	return p.repo.AddPremium(models.FPPremium{
		FPPremiumID:          uuid.New().String(),
		FPPremiumName:        premium.FPPremiumName,
		FPPremiumDescription: premium.FPPremiumDescription,
		FPPremiumServiceID:   premium.FPPremiumServiceID,
		FPPremiumStatus:      premium.FPPremiumStatus,
	})
}

func (p Premium) UpdatePremium(premium dto.FPUpdatePremium) (models.FPPremium, error) {

	fmt.Println(premium)

	isValidGuid := helper.IsValidUUID(premium.FPPremiumID)
	if !isValidGuid {
		return models.FPPremium{}, errors.New("invalid premiumId")
	}

	return p.repo.UpdatePremium(models.FPPremium{
		FPPremiumID:          premium.FPPremiumID,
		FPPremiumName:        premium.FPPremiumName,
		FPPremiumDescription: premium.FPPremiumDescription,
		FPPremiumServiceID:   premium.FPPremiumServiceID,
		FPPremiumStatus:      premium.FPPremiumStatus,
	})
}

func (p Premium) FindAllPremium() ([]models.FPPremium, error) {
	return p.repo.FindAllPremium()
}

func (p Premium) DeletePremium(premiumId string, spId string) error {
	isValidGuid := helper.IsValidUUID(premiumId)
	if !isValidGuid {
		return errors.New("invalid premiumId")
	}

	isValidGuid = helper.IsValidUUID(spId)
	if !isValidGuid {
		return errors.New("invalid serviceProviderId")
	}
	return p.repo.DeletePremium(premiumId, spId)
}

func (p Premium) FindPremium(premiumId string, spId string) (models.FPPremium, error) {
	isValidGuid := helper.IsValidUUID(premiumId)
	if !isValidGuid {
		return models.FPPremium{}, errors.New("invalid premiumId")
	}

	isValidGuid = helper.IsValidUUID(spId)
	if !isValidGuid {
		return models.FPPremium{}, errors.New("invalid serviceProviderId")
	}
	return p.repo.FindPremium(premiumId, spId)
}

func (p Premium) FindPremiumBySP(spId string) ([]models.FPPremium, error) {

	isValidGuid := helper.IsValidUUID(spId)
	if !isValidGuid {
		return []models.FPPremium{}, errors.New("invalid serviceProviderId")
	}
	return p.repo.FindPremiumBySP(spId)
}
