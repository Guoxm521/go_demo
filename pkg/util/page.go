package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetPage(c *gin.Context) int {
	result := 1
	page := com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		result = page
	}
	return result
}

func GetSize(c *gin.Context) int {
	result := 10
	size := com.StrTo(c.Query("size")).MustInt()
	if size > 0 {
		result = size
	}
	return result
}
