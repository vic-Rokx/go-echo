package routes

import (
	"errors"
	"example/api_gateway/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func DeleteUser(ctx echo.Context) error {
	id := ctx.Param("id")

	var user models.User

	result := models.DB.Delete(&user, "id = ?", id)

	if result.Error != nil {
		// Check if the error is because the records were not found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ctx.JSON(http.StatusNotFound, echo.Map{"error": result.Error})
		}
		// For any other error, return an internal server error
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete user"})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "User deleted", "Id": id})
}

func DeleteArticle(ctx echo.Context) error {
	id := ctx.Param("id")

	var article models.Article

	result := models.DB.Delete(&article, "id = ?", id)

	if result.Error != nil {
		// Check if the error is because the records were not found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ctx.JSON(http.StatusNotFound, echo.Map{"error": result.Error})
		}
		// For any other error, return an internal server error
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete article"})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "Article deleted", "Id": id})
}

// func WipeArticles(ctx echo.Context) error {
// 	var article models.Article

// 	result := models.DB.Delete(&article, "timestamp")
// }
