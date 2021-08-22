package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(migrate bool) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DBNAME"),
		os.Getenv("DB_PORT"),
	)

	// Set the correct log level according to the environment
	var logLevel int
	env := os.Getenv("GO_ENV")
	if env == "testing" || env == "production" {
		logLevel = int(logger.Error)
	} else {
		logLevel = int(logger.Info)
	}

	// Connect to the database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      logger.LogLevel(logLevel),
		}),
	})
	if err != nil {
		return err
	}

	if migrate {
		if err := automigrate(); err != nil {
			return err
		}
	}

	return nil
}

func automigrate() error {
	err := DB.AutoMigrate()

	return err
}

func Disconnect() error {
	// Acquire the raw DB object
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	// Close the DB connection
	err = sqlDB.Close()
	if err != nil {
		return err
	}

	return nil
}
