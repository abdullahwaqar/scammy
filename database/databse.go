package database

import (
	"api/database/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

// Bootstrap initilize the database and make a connection
func Bootstrap() (*gorm.DB, error) {
	dbConfig := os.Getenv("DATABASE_URL")
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
