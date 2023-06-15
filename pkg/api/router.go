package api

import "github.com/gin-gonic/gin"

//New defines api routes and returns *gin.Engine
func New() *gin.Engine {

	// set gin mode to release
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})
	return r
}
