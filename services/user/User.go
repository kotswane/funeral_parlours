package user

import (
	"funeral_parlour/models"
	"funeral_parlour/models/dto"
	"funeral_parlour/repository/user"
)

type User struct {
	svc user.UserRepository
}

func NewUser(svc user.UserRepository) UserService {
	return &User{svc: svc}
}

func (u User) AddUser(user dto.FPUser) (models.FPUser, error) {
	return u.svc.AddUser(models.FPUser{
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
	return u.svc.UpdateUser(models.FPUser{
		FPUserID:          user.FPUserID,
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
	return u.svc.FindAllUser()
}

func (u User) DeleteUser(userId int) error {
	return u.svc.DeleteUser(userId)
}

func (u User) FindUser(userId int) (models.FPUser, error) {

	return u.svc.FindUser(userId)
}
