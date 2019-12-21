package email

import (
	"api/database/models"

	common "api/lib"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Email type alias
type Email = models.Email

// JSON type alias
type JSON = common.JSON

func resgister(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		Email       string `json:"email" binding:"required"`
		EmailHeader string `json:"email_header" binding:"required"`
	}
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}
	email := Email{
		Email:               requestBody.Email,
		EmailHeader:         requestBody.EmailHeader,
		NumberOfOccurrences: 1,
	}
	db.NewRecord(email)
	db.Create(&email)
	c.JSON(200, email.Serialize())
}
