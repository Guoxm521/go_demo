package v1

import (
	"example.com/m/v2/models"
	"example.com/m/v2/pkg/app"
	"example.com/m/v2/pkg/e"
	"example.com/m/v2/service/spider_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MovieAdd struct {
	Title    string `form:"title" json:"title",des:"电影名称"`
	Time     string `form:"time" json:"time'",des:"上映时间"`
	Duration string `form:"duration" json:"duration",des:"播放时长"`
	Director string `form:"director" json:"director",des:"导演"`
	FilmType string `form:"filmType" json:"filmType",des:"类型"`
	Address  string `form:"address" json:"address",des:"制片国家/地区"`
	Language string `form:"language" json:"language",des:"语言"`
	Des      string `form:"des" json:"des",des:"描述"`
}

func DoubanMovieList(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	movie := models.DoubanMovie{
		Title: "标题",
	}
	res, err := movie.List()
	if err != nil {
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, res)
}

func DoubanMovieAdd(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	spider_service.DouBanMovie()
	appG.Response(http.StatusOK, e.SUCCESS, "添加")
}
