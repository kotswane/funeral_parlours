package database

import (
	"fmt"
	"funeral_parlour/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDataBase() (*gorm.DB, error) {
	dataBaseConfigs := config.LoadDataBaseConfigs()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dataBaseConfigs.DBUser, dataBaseConfigs.DBPassword, dataBaseConfigs.DBHost, dataBaseConfigs.DBPort, dataBaseConfigs.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
