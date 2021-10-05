package healthz

import "github.com/gin-gonic/gin"

func New() *gin.Engine {

	// set gin mode to release
	gin.SetMode(gin.ReleaseMode)

	h := gin.New()
	h.Use(gin.Recovery())
	h.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})
	return h
}
