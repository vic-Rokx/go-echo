package routes

import (
	"example/api_gateway/models"
	"example/api_gateway/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddArticle(ctx echo.Context) error {
	var input struct {
		Title             string `json:"title"`
		Author            string `json:"author"`
		Source            string `json:"source"`
		Image             string `json:"image"`
		Summary           string `json:"summary"`
		ArticleCategories []uint `gorm:"many2many:article_categories;"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"Error": err.Error()})
	}

	newArticle := models.Article{
		Title:   input.Title,
		Author:  input.Author,
		Source:  input.Source,
		Image:   input.Image,
		Summary: input.Summary,
	}

	var categories []models.Category
	if len(input.ArticleCategories) > 0 {
		if err := models.DB.Where("id IN ?", input.ArticleCategories).Find(&categories).Error; err != nil {
			return ctx.String(http.StatusInternalServerError, "Failed to find categories")
		}
	}

	newArticle.ArticleCategories = categories

	_, err := utils.ValidateArticle(newArticle)

	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"message": "Error with validation"})
	}

	if result := models.DB.Create(&newArticle); result.Error != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"Error": "Failed to create the new article"})
	}

	return ctx.JSON(http.StatusOK, newArticle)
}

func AddCategory(ctx echo.Context) error {
	var input struct {
		Name string `json:"name"`
	}

	// Bind the JSON payload to the input struct
	if err := ctx.Bind(&input); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	// Initialize the category with the provided input
	category := models.Category{
		Name: input.Name,
	}

	// Create the category in the database
	if result := models.DB.Create(&category); result.Error != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
	}

	// Return the ID of the created category
	return ctx.JSON(http.StatusOK, category.ID)
}
