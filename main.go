package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

var Db *gorm.DB

func initDB() {
	dsn := "host='localhost' user='postgres' dbname='postgres' password='123456' port=5432 sslmode=disable"
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Db", err)
	}
	Db.AutoMigrate(&Meme{})
}
func rateLimiter() gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Every(time.Minute), 5)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "To many requests, chill out"})
			return
		}
		c.Next()
	}
}
func getMeme(c *gin.Context) {
	var memes []Meme
	if result := Db.Find(&memes); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, memes)
}

func getMemeByID(c *gin.Context) {
	var meme Meme
	id := c.Param("id")
	if err := Db.Where("IDMeme = ?", id).First(&meme, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Meme not found"})
		return
	}
	c.JSON(http.StatusOK, meme)
}

func createMeme(c *gin.Context) {
	var NewMeme Meme
	if err := c.BindJSON(&NewMeme); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	Db.Create(&NewMeme)
	c.JSON(http.StatusCreated, NewMeme)
}

func updateMeme(c *gin.Context) {
	id := c.Param("id")
	var updatedData map[string]interface{}
	if err := c.BindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := Db.Model(&Meme{}).Where("IDMeme= ?", id).Updates(updatedData)
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

func deleteMeme(c *gin.Context) {
	id := c.Param("id")
	result := Db.Where("IDMeme = ?", id).Delete(&Meme{})
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Meme not found"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "meme deleted"})
}

func main() {
	initDB()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(rateLimiter())

	router.GET("/memes", getMeme)
	router.GET("/memes/:id", getMemeByID)
	router.POST("/memes", createMeme)
	router.PUT("/memes/:id", updateMeme)
	router.DELETE("/memes/:id", deleteMeme)
	router.Run(":8080")
}
