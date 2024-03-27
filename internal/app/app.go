package app

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"rams/pkg/models"
)

func InitConfigs() *models.AppConfigs {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Error("Error loading .env file")
	}

	database := &models.Postgres{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}

	return &models.AppConfigs{
		Port:     os.Getenv("PORT"),
		Database: database,
	}
}
