package main

import (
	"APIdinyPodkluchiny/internal"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db := internal.InitDB()
	err := internal.AutoMigrate(db)
	if err != nil {
		log.Fatal("Failed to migrate Db", err)
	}
	router := gin.Default()
	router.Use(internal.CORSMiddleware())
	router.Use(internal.RateLimiter())
	router.Use(internal.AuthMiddleware())

	router.GET("/memes", func(c *gin.Context) { internal.GetMeme(c, db) })
	router.GET("/memes/:id", func(c *gin.Context) { internal.GetMemeByID(c, db) })
	router.POST("/memes", func(c *gin.Context) { internal.CreateMeme(c, db) })
	router.PUT("/memes/:id", func(c *gin.Context) { internal.UpdateMeme(c, db) })
	router.DELETE("/memes/:id", func(c *gin.Context) { internal.DeleteMeme(c, db) })
	router.Run(":8080")
}
