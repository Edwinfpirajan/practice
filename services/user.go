package services

import (
	"github.com/Edwinfpirajan/practice.git/config"
	"github.com/Edwinfpirajan/practice.git/models"
)

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}
