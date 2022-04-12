package models

import (
	"encoding/json"
	"fmt"
)

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func AddTag(name string, state int, createdBy string) (interface{}, error) {
	tag := Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}
	result := db.Create(&tag)
	value := result.Value
	_b, _ := json.Marshal(value)
	fmt.Println("+++++++++++++++++", _b)
	var a = make(map[string]interface{}, 0)
	_ = json.Unmarshal(_b, &a)
	fmt.Println("aaaaaa", a)
	fmt.Println("bbbbbbb", a["id"])

	if result.Error != nil {
		return nil, result.Error
	}
	return result.Value, nil
}
