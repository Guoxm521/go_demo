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
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_TAGS_FAIL, err.Error())
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
	httpCode, errCode, err := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, err)
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

type DeleteTagForm struct {
	ID int `form:"id" json:"id" binding:"required"`
}

// DeleteTag 删除文章标签
func DeleteTag(c *gin.Context) {
	var (
		form DeleteTagForm
		appG = app.Gin{C: c}
	)
	httpCode, errCode, err := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, err.Error())
		return
	}
	tagService := tag_service.Tag{
		ID: form.ID,
	}
	exists, err := tagService.ExistTagByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}
	if err := tagService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_TAG_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, "success")
}
