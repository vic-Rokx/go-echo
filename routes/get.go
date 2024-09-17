package routes

import (
	"errors"
	"example/api_gateway/models"
	helpers "example/api_gateway/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetUserByID(ctx echo.Context) error {
	id := ctx.Param("id")
	user, err := helpers.GetUserByIdHelper((id))
	if err != nil {
		return ctx.String(http.StatusNotFound, "The user was not found")

	}
	return ctx.JSON(http.StatusOK, user)
}

func GetArticleById(ctx echo.Context) error {
	id := ctx.Param("id")
	var article models.Article

	result := models.DB.Preload("ArticleCategories").First(&article, "id = ?", id)

	if result.Error != nil {
		// Check if the error is because the records were not found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ctx.JSON(http.StatusNotFound, echo.Map{"error": result.Error})
		}
		// For any other error, return an internal server error
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve articles"})
	}

	return ctx.JSON(http.StatusOK, article)
}

func GetUsers(ctx echo.Context) error {
	var users []models.User

	result := models.DB.Preload("UserCategories").Find(&users)
	if result.Error != nil {
		// Check if the error is because the records were not found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ctx.JSON(http.StatusNotFound, echo.Map{"error": "No users found"})
		}
		// For any other error, return an internal server error
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve users"})
	}

	if len(users) == 0 {
		// If no users are found, it's also useful to return a not found status
		return ctx.JSON(http.StatusNotFound, echo.Map{"message": "No users found"})
	}

	return ctx.JSON(http.StatusOK, users)
}

func GetCategories(ctx echo.Context) error {
	var categories []models.Category

	result := models.DB.Find(&categories)
	if result.Error != nil {
		// Check if the error is because the records were not found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ctx.JSON(http.StatusNotFound, echo.Map{"error": "No categories found"})
		}
		// For any other error, return an internal server error
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve categories"})
	}

	if len(categories) == 0 {
		// If no categories are found, it's also useful to return a not found status
		return ctx.JSON(http.StatusNotFound, echo.Map{"message": "No categories found"})
	}

	return ctx.JSON(http.StatusOK, categories)
}

func GetArticles(ctx echo.Context) error {
	var articles []models.Article

	result := models.DB.Find(&articles)
	if result.Error != nil {
		// Check if the error is because the records were not found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ctx.JSON(http.StatusNotFound, echo.Map{"error": "No articles found"})
		}
		// For any other error, return an internal server error
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve articles"})
	}

	if len(articles) == 0 {
		// If no categories are found, it's also useful to return a not found status
		return ctx.JSON(http.StatusNotFound, echo.Map{"message": "No articles found"})
	}

	return ctx.JSON(http.StatusOK, articles)
}

func GetArticlesByCategories(ctx echo.Context) error {
	category := ctx.Param("category")
	news, err := helpers.GetNewsByCategoryHelper((category))
	if err != nil {
		return ctx.String(http.StatusNotFound, "The News category was not found")
	}

	return ctx.JSON(http.StatusOK, news)
}
