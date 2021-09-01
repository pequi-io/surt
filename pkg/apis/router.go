package apis

import "github.com/gin-gonic/gin"

// SetupRouter defines api routes and returns Engine
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
