package config

import "os"

type DataBase struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBPort     string
	DBHost     string
	DBDriver   string
}

func LoadDataBaseConfigs() DataBase {
	return DataBase{
		DBName:     os.Getenv("DBName"),
		DBUser:     os.Getenv("DBUser"),
		DBPassword: os.Getenv("DBPassword"),
		DBPort:     os.Getenv("DBPort"),
		DBHost:     os.Getenv("DBHost"),
		DBDriver:   os.Getenv("DBDriver"),
	}
}
