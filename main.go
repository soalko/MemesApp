package main

import (
	"APIdinyPodkluchiny/DB"
	"APIdinyPodkluchiny/api"
	_ "APIdinyPodkluchiny/docs" // docs генерируются Swagger
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

//	@title			MemesApp
//	@version		0.2
//	@description	Система пояснения сложных информационных сообщений юмористического характера.
//	@termsOfService	http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/

// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse

func main() {
	db := DB.InitDB()
	err := DB.AutoMigrate(db)
	if err != nil {
		log.Fatal("Failed to migrate Db", err)
	}
	router := gin.Default()
	router.Use(api.CORSMiddleware())
	router.Use(api.RateLimiter())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/memes", func(c *gin.Context) { DB.GetMeme(c, db) })
	router.GET("/memes/:id", func(c *gin.Context) { DB.GetMemeByID(c, db) })
	router.POST("/memes", func(c *gin.Context) { DB.CreateMeme(c, db) })
	router.PUT("/memes/:id", func(c *gin.Context) { DB.UpdateMeme(c, db) })
	router.DELETE("/memes/:id", func(c *gin.Context) { DB.DeleteMeme(c, db) })
	router.Run(":8080")
}
