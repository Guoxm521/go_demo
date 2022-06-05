package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
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
	//value := result.Value
	//_b, _ := json.Marshal(value)
	//var a = make(map[string]interface{}, 0)
	//_ = json.Unmarshal(_b, &a)
	//fmt.Println("反射获取对象", a)
	//fmt.Println("反射获取对象", a["id"])
	if result.Error != nil {
		return nil, result.Error
	}
	return result.Value, nil
}

func GetTags(pageNum int, pageSize int, maps interface{}) ([]Tag, error) {
	var (
		tags []Tag
		err  error
	)
	if pageSize > 0 && pageNum > 0 {
		fmt.Println("pageNum", pageNum, pageSize)
		err = db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags).Error
	}
	return tags, err
}

func ExistTagByID(id int) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ? AND deleted_on = ?", id, 0).First(&tag).Error
	defer db.Close()
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

func DeleteTag(id int) error {
	if err := db.Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}
	return nil
}

func EditTag(id int, data interface{}) error {
	err := db.Model(&Tag{}).Where("id = ? AND deleted_on = ?", id, 0).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}
