package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// GetTags 获取多个标签
func GetTags(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "获取标签",
	})
}

type AddTagForm struct {
	Name      string `form:"name" valid:"Required;MaxSize(100)"`
	CreatedBy string `form:"created_by" valid:"Required;MaxSize(100)"`
	State     int    `form:"state" valid:"Range(0,1)"`
}

// AddTag 新增文章标签
func AddTag(c *gin.Context) {
	var (
		form AddTagForm
	)
	err := c.Bind(&form)
	if err != nil {
		fmt.Println("解析失败")
		return
	}
	fmt.Printf("%v", form)
	c.JSON(200, gin.H{
		"message": "添加标签",
	})
}

// EditTag 编辑文章标签
func EditTag(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "编辑标签",
	})
}

// DeleteTag 删除文章标签
func DeleteTag(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "删除标签",
	})
}
