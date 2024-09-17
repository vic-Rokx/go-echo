package utils

import (
	"example/api_gateway/models"
	"fmt"
)

func ValidateArticle(newArticle models.Article) (bool, error) {
	article := newArticle
	err := models.ValidateArticle(article)

	if err != nil {
		fmt.Println("Validation failed:", err)
		return false, err
	} else {
		fmt.Println("Validation successful!")
		return true, nil
	}
}
