package DB

import (
	"APIdinyPodkluchiny/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetMeme(c *gin.Context, db *gorm.DB) {
	var memes []models.Meme
	if err := db.Find(&memes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, memes)
}

func GetMemeByID(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var meme models.Meme

	if err := db.First(&meme, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Meme not found"})
		return
	}
	c.JSON(http.StatusOK, meme)
}

func CreateMeme(c *gin.Context, db *gorm.DB) {
	var NewMeme models.Meme
	if err := c.ShouldBindJSON(&NewMeme); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&NewMeme).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, NewMeme)
}

func UpdateMeme(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var updatedData map[string]interface{}
	if err := c.BindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Model(&models.Meme{}).Where("idmeme= ?", id).Updates(updatedData)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Meme not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Meme updated"})
}

func DeleteMeme(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	result := db.Where("idmeme = ?", id).Delete(&models.Meme{})
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Meme not found"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "meme deleted"})
}
