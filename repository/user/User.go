package user

import (
	"funeral_parlour/models"
	"funeral_parlour/models/dto"

	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) UserRepository {
	return &User{db: db}
}

func (u User) UpdatePassword(password dto.FPPassword) error {

	if err := u.db.Model(&models.FPUser{}).Where("usr_id = ?", password.UserID).Update("usr_password", password.Password).Error; err != nil {
		return err
	}
	return nil
}

func (u User) AddUser(user models.FPUser) (models.FPUser, error) {
	results := u.db.Create(user)
	if results.Error != nil {
		return models.FPUser{}, results.Error
	}
	return user, nil
}

func (u User) UpdateUser(user models.FPUser, id string) (models.FPUser, error) {
	if err := u.db.Model(&models.FPUser{}).Where("usr_id = ?", id).Updates(user).Error; err != nil {
		return models.FPUser{}, err
	}
	return user, nil
}

func (u User) DeleteUser(userId string) error {
	if err := u.db.Where("usr_id = ?", userId).Delete(&models.FPUser{}).Error; err != nil {
		return err
	}
	return nil
}

func (u User) FindAllUser() ([]models.FPUser, error) {
	var userList []models.FPUser
	if err := u.db.Find(&userList).Error; err != nil {
		return userList, err
	}
	return userList, nil
}

func (u User) FindUser(userId string) (models.FPUser, error) {
	var user models.FPUser
	if err := u.db.Where("usr_id = ?", userId).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
