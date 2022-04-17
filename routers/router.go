package routers

import (
	v1 "example.com/m/v2/routers/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//gin.SetMode(setting.RunMode)
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test index",
		})
	})

	apiv1 := r.Group("/api")
	{
		apiv1.GET("/tags/list", v1.GetTagsList)
		apiv1.POST("/tags/add", v1.AddTag)
		apiv1.POST("/tags/edit", v1.EditTag)
		apiv1.GET("/tags/del", v1.DeleteTag)
	}
	return r
}
