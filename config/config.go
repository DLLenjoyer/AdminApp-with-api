package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func SetupDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=user password=password dbname=mydb port=5432 sllmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Подключение к базе данных успешно!")
	return db, nil
}
