package route

import (
	"document-server/configs"
	"github.com/gin-gonic/gin"
)

// RegisterRouter Configure route forwarding
func RegisterRouter() *gin.Engine {
	r := gin.Default()
	gv1 := r.Group("v1")
	gv1.Use()
	{
		gv1.GET("/get/config", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": configs.Conf,
			})
		})
	}
	return r
}
