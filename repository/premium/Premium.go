package premium

import (
	"funeral_parlour/models"

	"gorm.io/gorm"
)

type Premium struct {
	db *gorm.DB
}

func NewPremium(db *gorm.DB) PremiumRepository {
	return &Premium{db: db}
}

func (p Premium) AddPremium(premium models.FPPremium) (models.FPPremium, error) {
	response := p.db.Create(premium)
	if response.Error != nil {
		return models.FPPremium{}, response.Error
	}
	return premium, nil
}

func (p Premium) UpdatePremium(premium models.FPPremium) (models.FPPremium, error) {
	updatePremium := models.FPPremium{
		FPPremiumName:        premium.FPPremiumName,
		FPPremiumDescription: premium.FPPremiumDescription,
		FPPremiumStatus:      premium.FPPremiumStatus,
	}
	if err := p.db.Model(&models.FPPremium{}).Where("fp_premium_id = ? AND fp_premium_sp_id = ?", premium.FPPremiumID, premium.FPPremiumServiceID).Updates(updatePremium).Error; err != nil {
		return models.FPPremium{}, err
	}
	return premium, nil
}

func (p Premium) FindAllPremium() ([]models.FPPremium, error) {
	var premiumList []models.FPPremium
	if err := p.db.Find(&premiumList).Error; err != nil {
		return nil, err
	}
	return premiumList, nil
}

func (p Premium) DeletePremium(premiumId string, spId string) error {
	if err := p.db.Where("fp_premium_id = ? AND fp_premium_sp_id = ?", premiumId, spId).Delete(&models.FPPremium{}).Error; err != nil {
		return err
	}
	return nil
}

func (p Premium) FindPremium(premiumId string, spId string) (models.FPPremium, error) {
	var premium models.FPPremium
	if err := p.db.Where("fp_premium_id = ? AND fp_premium_sp_id = ?", premiumId, spId).First(&premium).Error; err != nil {
		return premium, err
	}
	return premium, nil
}

func (p Premium) FindPremiumBySP(spId string) ([]models.FPPremium, error) {
	var premium []models.FPPremium
	if err := p.db.Where("fp_premium_sp_id = ?", spId).Find(&premium).Error; err != nil {
		return premium, err
	}
	return premium, nil
}
