package main

import (
	"APIdinyPodkluchiny/DB"
	"APIdinyPodkluchiny/api"
	_ "APIdinyPodkluchiny/docs" // docs генерируются Swagger
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

//	@title			MemesApp
//	@version		0.3
//	@description	Система пояснения сложных информационных сообщений юмористического характера.
//	@termsOfService	http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		178.250.158.143:8080
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
	router.Use(func(c *gin.Context) { api.SaveUsrLogs(c) })
	fmt.Println("AUTH_ENABLED:", os.Getenv("AUTH_ENABLED"))
	public := router.Group("/")
	{
		public.GET("/health", func(c *gin.Context) { DB.PingDB(c) })
		public.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		router.GET("/memes", func(c *gin.Context) { DB.GetMeme(c, db) })
		router.GET("/memes/:id", func(c *gin.Context) { DB.GetMemeByID(c, db) })
		router.POST("/memes", func(c *gin.Context) { DB.CreateMeme(c, db) })
		router.PATCH("/memes/:id", func(c *gin.Context) { DB.UpdateMeme(c, db) })
		router.DELETE("/memes/:id", func(c *gin.Context) { DB.DeleteMeme(c, db) })
		router.POST("/tags", func(c *gin.Context) { DB.CreateTag(c, db) })
		router.PATCH("/tags/:id", func(c *gin.Context) { DB.UpdateTag(c, db) })
		router.DELETE("/tags/:id", func(c *gin.Context) { DB.DeleteTag(c, db) })
		router.POST("/categories", func(c *gin.Context) { DB.CreateCategory(c, db) })
		router.PATCH("/categories/:id", func(c *gin.Context) { DB.UpdateCategory(c, db) })
		router.DELETE("/categories/:id", func(c *gin.Context) { DB.DeleteCategory(c, db) })
	}
	private := router.Group("/")
	private.Use(api.AuthMiddleware())
	{
		//пока для теста всё public
	}
	router.Run(":8080")
}
