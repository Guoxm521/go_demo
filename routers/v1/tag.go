package v1

import (
	"example.com/m/v2/pkg/app"
	"example.com/m/v2/pkg/e"
	"example.com/m/v2/pkg/util"
	"example.com/m/v2/service/tag_service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetTagsList 获取列表
func GetTagsList(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	name := c.Query("name")
	state := -1
	if arg := c.Query("state"); arg != "" {
		int, _ := strconv.Atoi(arg)
		state = int
	}
	tagService := tag_service.Tag{
		Name:     name,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: util.GetSize(c),
	}
	tags, err := tagService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_TAGS_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": tags,
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
