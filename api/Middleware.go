package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var jwtKey = []byte("my_secret_key")

func SaveUsrLogs(c *gin.Context) {
	log.Printf("Request: %s %s", c.Request.Method, c.Request.URL)
	c.Next()
}
func RateLimiter() gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Every(time.Minute), 50)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "To many requests, chill out"})
			return
		}
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if os.Getenv("AUTH_ENABLED") != "false" {
			authHeader := c.GetHeader("Authorization")
			if authHeader == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is empty"})
				return
			}
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})
			if err != nil || !token.Valid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
				return
			}
			c.Next()
		} else {
			c.Next()
		}

	}
}

/*
	ОТКЛЮЧИТЬ ДОЛБАННУЮ АВТОРИЗАЦИЮ ДЛЯ ТЕСТА ЧЕРЕЗ swagger!
	if strings.HasPrefix(c.Request.URL.Path, "/swagger") ||
		strings.HasPrefix(c.Request.URL.Path, "/docs") ||
		c.Request.URL.Path == "/v2/api-docs" {
		c.Next()
		return
	}
*/
