package main

import (
	"api/api"
	"api/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
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
	app.Use(cors.Default()) // allows all origins
	api.ApplyRoutes(app)    // apply api router
	app.Run(":" + port)     // listen to given port
}
