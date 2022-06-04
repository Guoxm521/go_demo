package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func Console(data interface{}) {
	bs, err := json.Marshal(data)
	if err != nil {
		fmt.Println("解析失败")
	}
	var out bytes.Buffer
	json.Indent(&out, bs, "", "\t")
	fmt.Printf("%v\n", out.String())
}

// DirExits 判断文件夹是否存在
func DirExits(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// DirCreate 生成目录文件夹
func DirCreate(_dir string) error {
	exist, err := DirExits(_dir)
	if err != nil {
		return errors.New("get dir error! " + err.Error())
	}
	if !exist {
		err := os.MkdirAll(_dir, os.ModePerm)
		if err != nil {
			return errors.New("mkdir failed! " + err.Error())
		}
	}
	return nil
}
