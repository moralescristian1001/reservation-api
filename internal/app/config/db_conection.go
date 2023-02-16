package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {

	dsn := fmt.Sprintf(
		"host=wispy-flower-5019-db.internal user=postgres password=bCUUSR1CtixNmgB dbname=reservation port=5432 sslmode=disable",
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to the database.")

	return db
}