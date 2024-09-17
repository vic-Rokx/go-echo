package utils

import (
	"errors"
	"example/api_gateway/models"

	"gorm.io/gorm"
)

func GetUserByIdHelper(id string) (*models.User, error) {

	var user models.User

	// Assuming `db` is your *gorm.DB instance
	result := models.DB.Preload("UserCategories").First(&user, "id = ?", id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("User not found")
		} else {
			return nil, errors.New("User not found")
		}
	}

	// Responding with the user data
	return &user, nil
}

func GetNewsByCategoryHelper(category string) ([]*models.Article, error) {
	return nil, errors.New("Category not found")
}
