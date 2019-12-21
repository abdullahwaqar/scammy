package database

import (
	"api/database/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

// Initilize the database and make a connection
func Bootstrap() (*gorm.DB, error) {
	dbConfig := os.Getenv("DB_CONFIG")
	db, err := gorm.Open("postgres", dbConfig)
	db.LogMode(true) //* Logs SQL
	if err != nil {
		panic(err)
	}
	fmt.Println("Conneted to database")
	//* Do the migrations here
	models.Migrate(db)
	return db, err
}
