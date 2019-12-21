package email

import "github.com/gin-gonic/gin"

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	emails := r.Group("/")
	{
		emails.GET("/all", getAll)        //* Get all scammers emails in databse
		emails.POST("/newscam", register) //* Posts new scammer email
		emails.POST("/scan", scan)        //* Scans an email in the database
		emails.PATCH("/:id", update)
	}
}
