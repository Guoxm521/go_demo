package v1

import (
	"example.com/m/v2/pkg/app"
	"example.com/m/v2/pkg/e"
	"example.com/m/v2/service/tag_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetTags 获取多个标签
func GetTags(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "获取标签",
	})
}

type AddTagForm struct {
	Name      string `form:"name" json:"name" binding:"required,max=10"`
	CreatedBy string `form:"created_by" json:"created_by" binding:"required"`
	State     int    `form:"state" json:"state" binding:"oneof=5 7 9"`
}

// AddTag 新增文章标签
func AddTag(c *gin.Context) {
	var (
		form AddTagForm
		appG = app.Gin{C: c}
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
	}
	tagService := tag_service.Tag{
		Name:      form.Name,
		CreatedBy: form.CreatedBy,
		State:     form.State,
	}
	res, err := tagService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, res)
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
