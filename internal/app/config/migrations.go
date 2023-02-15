package config

import (
	"fmt"
	"gorm.io/gorm"
	"reservation-api/internal/command/infraestructure/storage/dto"
)

func ExecuteMigrations(db *gorm.DB)  {

	err := db.AutoMigrate(&dto.Reservation{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Migrations were successfully executed.")
}
