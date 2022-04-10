package util

import (
	"bytes"
	"encoding/json"
	"fmt"
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
