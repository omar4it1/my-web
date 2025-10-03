package main

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/news", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"main":   "Morocco's Gen Z protests: What you need to know- this is added on server",
			"author": "Mohamed Moustaid in Rabat",
			"sub":    "After almost a week of regular demonstrations, protests in Morocco turned violent, resulting in hundreds injured and two deaths. Who's behind the protests, what are they demanding, and where might this end?",
		})
	})
	router.Run()
}
