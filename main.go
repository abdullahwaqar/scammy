package main

import (
	"api/api"
	"api/database"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load .env environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// initializes database
	db, _ := database.Bootstrap()
	port := os.Getenv("PORT")
	app := gin.Default() // create gin app
	app.Use(database.Inject(db))
	api.ApplyRoutes(app) // apply api router
	app.Run(":" + port)  // listen to given port
}
