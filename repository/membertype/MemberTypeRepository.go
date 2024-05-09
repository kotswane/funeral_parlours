package membertype

import (
	"funeral_parlour/models"
)

type MemberTypeRepository interface {
	AddMemberType(membertype models.FPMemberType) (models.FPMemberType, error)
	UpdateMemberType(membertype models.FPMemberType) (models.FPMemberType, error)
	FindAllMemberType() ([]models.FPMemberType, error)
	DeleteMemberType(membertypeId string) error
	FindMemberType(membertypeId string) (models.FPMemberType, error)
}
