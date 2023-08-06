package migrations

import (
	"github.com/Edwinfpirajan/practice.git/config"
	"github.com/Edwinfpirajan/practice.git/models"
)

func GenerateMigration() {
	config.DB.AutoMigrate(&models.User{}, &models.Team{})
}
