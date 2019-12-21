package apiv1

import "github.com/gin-gonic/gin"

import "api/api/v1.0/email"

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "I am alive!",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		v1.GET("/ping", ping)
		email.ApplyRoutes(v1)
	}
}
