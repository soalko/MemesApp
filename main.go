package main

import (
	"APIdinyPodkluchiny/DB"
	"APIdinyPodkluchiny/api"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db := DB.InitDB()
	err := DB.AutoMigrate(db)
	if err != nil {
		log.Fatal("Failed to migrate Db", err)
	}
	router := gin.Default()
	router.Use(api.CORSMiddleware())
	router.Use(api.RateLimiter())

	router.GET("/memes", func(c *gin.Context) { DB.GetMeme(c, db) })
	router.GET("/memes/:id", func(c *gin.Context) { DB.GetMemeByID(c, db) })
	router.POST("/memes", func(c *gin.Context) { DB.CreateMeme(c, db) })
	router.PUT("/memes/:id", func(c *gin.Context) { DB.UpdateMeme(c, db) })
	router.DELETE("/memes/:id", func(c *gin.Context) { DB.DeleteMeme(c, db) })
	router.Run(":8080")
}
