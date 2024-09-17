package controllers

import (
	"errors"
	"example/api_gateway/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SignUp(ctx echo.Context) error {
	var input struct {
		Name           string `json:"name"`
		Email          string `json:"email"`
		Password       string `json:"password"`
		UserCategories []struct {
			Name string `json:"name"`
		} `json:"user_categories"`
	}

	// Bind the JSON payload to the input struct
	if err := ctx.Bind(&input); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)

	if err != nil {
		log.Fatal("Error generating password")
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	// Initialize the user with the provided input, except for NewsCategories
	user := models.User{
		ID:       uuid.NewString(),
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hash), // Remember to hash the password in a real application
	}

	var categories []models.Category
	if len(input.UserCategories) > 0 {
		var categoryNames []string
		for _, category := range input.UserCategories {
			categoryNames = append(categoryNames, category.Name)
		}
		// Use categoryNames in the query, not input.UserCategories
		if err := models.DB.Where("name IN ?", categoryNames).Find(&categories).Error; err != nil {
			return ctx.String(http.StatusInternalServerError, "Failed to find categories")
		}
	}
	user.UserCategories = categories

	// Create the user in the database with their categories
	if result := models.DB.Create(&user); result.Error != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to create user")

	}

	return ctx.JSON(http.StatusOK, user.ID)
}

func Login(ctx echo.Context) error {

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to find user")
	}

	var user models.User

	var result = models.DB.First(&user, "email = ?", input.Email)

	// Assuming `db` is your *gorm.DB instance

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		} else {
			return errors.New("user not found")
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Password is incorrect"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to create token")
	}

	cookie := new(http.Cookie)
	cookie.Name = "Authorization"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour * 30)
	ctx.SetCookie(cookie)
	return ctx.String(http.StatusOK, "write a cookie")

	// return ctx.JSON(http.StatusOK, echo.Map{"token": tokenString})

}

func Validate(ctx echo.Context) error {
	user := ctx.Get("user")

	return ctx.JSON(http.StatusOK, user)
}
