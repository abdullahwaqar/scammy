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

/*
* @func
*		Get all the emails form db, LIMIT(100)
 */
func getAll(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	cursor := c.Query("cursor")

	var emails []Email
	if cursor == "" {
		if err := db.Find(&emails).Limit(100).Order("no_of_occurrences").Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	}

	length := len(emails)
	serialized := make([]JSON, length, length)
	for i := 0; i < length; i++ {
		serialized[i] = emails[i].Serialize()
	}
	c.JSON(200, serialized)
}

/*
* @func
*		Reads the email in the request body and searches it up in the database
 */
func scan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		Email string `json:"email" binding:"required"`
	}
	var requestBody RequestBody
	var email Email
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}
	// auto preloads the related model
	// http://gorm.io/docs/preload.html#Auto-Preloading
	if err := db.Set("gorm:auto_preload", true).Where("email = ?", requestBody.Email).First(&email).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}
	c.JSON(200, email.Serialize())
}

/*
* @func
*		Route for posting new email in database
 */
func register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		Email       string `json:"email" binding:"required"`
		EmailHeader string `json:"email_header"`
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

func update(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	type RequestBody struct {
		vote string `json:"vote" binding:"required"`
	}

	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}

	var email Email
	if err := db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&email).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	email.NumberOfOccurrences = email.NumberOfOccurrences + 1
	db.Save(&email)
	c.JSON(200, email.Serialize())
}
