package membertype

import (
	"funeral_parlour/models"

	"gorm.io/gorm"
)

type MemberType struct {
	db *gorm.DB
}

func NewMemberType(db *gorm.DB) MemberTypeRepository {
	return &MemberType{db: db}
}

// AddMemberType implements MemberTypeRepository.
func (m *MemberType) AddMemberType(membertype models.FPMemberType) (models.FPMemberType, error) {
	response := m.db.Create(membertype)
	if response.Error != nil {
		return models.FPMemberType{}, response.Error
	}
	return membertype, nil
}

// DeleteMemberType implements MemberTypeRepository.
func (m *MemberType) DeleteMemberType(membertypeId string) error {
	if err := m.db.Where("fp_member_type_id = ?", membertypeId).Delete(&models.FPMemberType{}).Error; err != nil {
		return err
	}
	return nil
}

// FindAllMemberType implements MemberTypeRepository.
func (m *MemberType) FindAllMemberType() ([]models.FPMemberType, error) {
	var memberTypeList []models.FPMemberType
	if err := m.db.Find(&memberTypeList).Error; err != nil {
		return nil, err
	}
	return memberTypeList, nil
}

// FindMemberType implements MemberTypeRepository.
func (m *MemberType) FindMemberType(membertypeId string) (models.FPMemberType, error) {

	var memberType models.FPMemberType
	if err := m.db.Where("fp_member_type_id = ?", membertypeId).First(&memberType).Error; err != nil {
		return memberType, err
	}
	return memberType, nil
}

// UpdateMemberType implements MemberTypeRepository.
func (m *MemberType) UpdateMemberType(membertype models.FPMemberType) (models.FPMemberType, error) {

	updateMemberType := models.FPMemberType{
		FPMemberTypeDescription: membertype.FPMemberTypeDescription,
	}

	if err := m.db.Model(&models.FPMemberType{}).Where("fp_member_type_id = ?", membertype.FPMemberTypeID).Updates(updateMemberType).Error; err != nil {
		return models.FPMemberType{}, err
	}
	return membertype, nil
}
