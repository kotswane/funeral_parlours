package premium

import (
	"funeral_parlour/models"
)

type PremiumRepository interface {
	AddPremium(premium models.FPPremium) (models.FPPremium, error)
	UpdatePremium(premium models.FPPremium) (models.FPPremium, error)
	FindAllPremium() ([]models.FPPremium, error)
	DeletePremium(premiumId string, spId string) error
	FindPremium(premiumId string, spId string) (models.FPPremium, error)

	FindPremiumBySP(spId string) ([]models.FPPremium, error)
}
