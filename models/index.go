package models

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

var DB *gorm.DB

// netstat -vanp tcp | grep {port}
// to show all current connections

// sudo kill <PID>

// For uuid generaterion connect via psql and do CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

func InitDB() {

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&Category{}); err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&Article{}); err != nil {
		panic(err)
	}

	fmt.Println("Migrated database")

	DB = db

}
