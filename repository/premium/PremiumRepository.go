package premium

import (
	"funeral_parlour/models"
)

type PremiumRepository interface {
	AddPremium(premium models.FPPremium) (models.FPPremium, error)
	UpdatePremium(premium models.FPPremium) (models.FPPremium, error)
	FindAllPremium() ([]models.FPPremium, error)
	DeletePremium(premiumId int) error
	FindPremium(premiumId int) (models.FPPremium, error)
}
