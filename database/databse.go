package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//* Initilize the database and make a connection
func Bootstrap() (*gorm.DB, error) {
	dbConfig := os.Getenv("DB_CONFIG")
	db, err := gorm.Open("postgres", dbConfig)
	db.LogMode(true) //* Logs SQL
	if err != nil {
		panic(err)
	}
	fmt.Println(("Conneted to database"))
	//* Do the migrations here
	// models.Mi
}
