package routes

import (
	"example/api_gateway/models"
	"fmt"
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

func autoMigrate() error {
	if err := models.DB.AutoMigrate(&models.User{}); err != nil {
		panic(err)
	}

	if err := models.DB.AutoMigrate(&models.Category{}); err != nil {
		panic(err)
	}

	if err := models.DB.AutoMigrate(&models.Article{}); err != nil {
		panic(err)
	}

	fmt.Println("Auto Migrated database")

	return nil
}

func Wipe(ctx echo.Context) error {
	// Convert []string to []interface{}
	tablesInterface := make([]interface{}, len(models.Tables))
	for i, v := range models.Tables {
		tablesInterface[i] = strings.ToLower(v)
	}

	if err := models.DB.Migrator().DropTable("categories"); err != nil {
		fmt.Println("Failed to drop")
		tablesNames := strings.Join(models.Tables, ", ")
		log.Fatalf("Failed to drop tables: %s. Error: %v", tablesNames, err)
	}

	fmt.Printf("Dropped tables: %s", tablesInterface)

	return nil
}
