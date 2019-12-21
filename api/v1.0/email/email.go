package email

import "github.com/gin-gonic/gin"

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	posts := r.Group("/posts")
	{
		posts.POST("/", resgister)
		// posts.GET("/", list)
		// posts.GET("/:id", read)
		// posts.DELETE("/:id", middlewares.Authorized, remove)
		// posts.PATCH("/:id", middlewares.Authorized, update)
	}
}
