package database

import (
	"fmt"
	"test-gorilla-mux/models"
	"test-gorilla-mux/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.Customer{},
		&models.Family{},
		&models.Nationality{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
