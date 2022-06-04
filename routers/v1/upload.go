package v1

import (
	"example.com/m/v2/pkg/app"
	"example.com/m/v2/pkg/e"
	"example.com/m/v2/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

func Upload(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	file, err := c.FormFile("file")
	//imagePath := c.PostForm("image_path")
	//size := c.PostForm("size")
	if err != nil {
		appG.Response(http.StatusOK, e.EMPTY_FILE, nil)
		return
	}
	//类型判断
	_fileType := file.Header.Get("Content-Type")
	_fileTypeList := []string{"image/jpeg", "image/png", "image/jpg", "image/gif"}
	isExit := false
	for _, s := range _fileTypeList {
		if _fileType == s {
			isExit = true
		}
	}
	if !isExit {
		appG.Response(http.StatusOK, e.FILE_TYPE_ERROR, nil)
		return
	}
	//文件大小
	_sizeMax := int64(2 * 1024 * 1024)
	_fileSize := file.Size
	if _fileSize > _sizeMax {
		appG.Response(http.StatusOK, e.OUT_FILE_SIZE, map[string]interface{}{
			"max": _sizeMax,
		})
		return
	}

	//保存
	fileExt := strings.ToLower(path.Ext(file.Filename))
	_fileName := fmt.Sprintf("%s%s", strconv.FormatInt(time.Now().UnixMicro(), 10), fileExt)
	_rootDir := "upload/"
	_fileNameSaveSlice := []string{"images"}
	fileDir := fmt.Sprintf("%s/%s/", strings.Join(_fileNameSaveSlice, "/"), time.Now().Format("20060102"))
	if _err := util.DirCreate(_rootDir + fileDir); _err != nil {
		appG.Response(http.StatusOK, e.ERROR, "")
		return
	}
	_imgPath := fmt.Sprintf("%s%s", fileDir, _fileName)
	//保存
	if err := c.SaveUploadedFile(file, _rootDir+_imgPath); err != nil {
		appG.Response(http.StatusOK, e.ERROR, "文件保存失败")
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"imgPath": fmt.Sprintf("%s/%s", _rootDir, _imgPath),
	})
}
