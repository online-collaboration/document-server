package route

import (
	"github.com/gin-gonic/gin"
)

// RegisterRouter Configure route forwarding
func RegisterRouter() *gin.Engine {
	r := gin.Default()
	gv1 := r.Group("v1")
	gv1.Use()
	{
		gv1.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "111",
			})
		})
		gv1.GET("/get/config", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "dddddd",
			})
		})
	}
	return r
}
