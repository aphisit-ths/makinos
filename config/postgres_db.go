package config

import (
	"fmt"
	"makino/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToPostgresDB(setting models.AppSettings) *gorm.DB {
	dsn := setting.ConnectionStrings.PostgresDb
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect PostgresDb ")
	}
	fmt.Println("### Connected to PostgresDb Successfully")

	return db
}
