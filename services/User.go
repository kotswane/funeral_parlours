package services

import (
	"errors"
	"funeral_parlour/models"
	"funeral_parlour/models/dto"
	"funeral_parlour/repository/user"
	"funeral_parlour/services/helper"

	"github.com/google/uuid"
)

type User struct {
	repo user.UserRepository
}

func NewUserService(repo user.UserRepository) UserService {
	return &User{repo: repo}
}

func (u User) AddUser(user dto.FPUser) (models.FPUser, error) {
	return u.repo.AddUser(models.FPUser{
		FPUserID:          uuid.NewString(),
		FPUserFirstName:   user.FPUserFirstName,
		FPUserMiddleName:  user.FPUserMiddleName,
		FPUserSurname:     user.FPUserSurname,
		FPUserIDNumber:    user.FPUserIDNumber,
		FPUserAddress:     user.FPUserAddress,
		FPUserPhoneNumber: user.FPUserPhoneNumber,
		FPUserEmail:       user.FPUserEmail,
		FPUserPassword:    user.FPUserPassword,
		FPUserRole:        user.FPUserRole,
		FPUserStatus:      user.FPUserStatus,
		FPUserSpID:        user.FPUserSpID,
		FPUSERUsername:    user.FPUSERUsername,
	})
}

func (u User) UpdateUser(user dto.FPUser) (models.FPUser, error) {
	isValidGuid := helper.IsValidUUID(user.FPUserID)
	if !isValidGuid {
		return models.FPUser{}, errors.New("invalid userId")
	}

	return u.repo.UpdateUser(models.FPUser{
		FPUserFirstName:   user.FPUserFirstName,
		FPUserMiddleName:  user.FPUserMiddleName,
		FPUserSurname:     user.FPUserSurname,
		FPUserIDNumber:    user.FPUserIDNumber,
		FPUserAddress:     user.FPUserAddress,
		FPUserPhoneNumber: user.FPUserPhoneNumber,
		FPUserEmail:       user.FPUserEmail,
		FPUserRole:        user.FPUserRole,
		FPUserStatus:      user.FPUserStatus,
		FPUserSpID:        user.FPUserSpID,
		FPUSERUsername:    user.FPUSERUsername,
	}, user.FPUserID)
}

func (u User) FindAllUser() ([]models.FPUser, error) {
	return u.repo.FindAllUser()
}

func (u User) DeleteUser(userId string) error {
	isValidGuid := helper.IsValidUUID(userId)
	if !isValidGuid {
		return errors.New("invalid userId")
	}
	return u.repo.DeleteUser(userId)
}

func (u User) FindUser(userId string) (models.FPUser, error) {
	isValidGuid := helper.IsValidUUID(userId)
	if !isValidGuid {
		return models.FPUser{}, errors.New("invalid userId")
	}
	return u.repo.FindUser(userId)
}
