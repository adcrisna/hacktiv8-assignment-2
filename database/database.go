package database

import (
	"assignment-2/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	var host = "localhost"
	var port = 5432
	var username = "postgres"
	var password = "legenda485132"
	var dbName = "db-go-sql"

	var conn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbName)

	db, err := gorm.Open(postgres.Open(conn))
	if err != nil {
		fmt.Println("error open connection to db", err)
	}
	db.Debug().AutoMigrate(models.Order{}, models.Item{})
	return db
}
