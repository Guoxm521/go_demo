package tag_service

import (
	"example.com/m/v2/models"
	"example.com/m/v2/service/cache_service"
)

type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int
	PageNum    int
	PageSize   int
}

func (t *Tag) Add() (interface{}, error) {
	return models.AddTag(t.Name, t.State, t.CreatedBy)
}

func (t *Tag) GetAll() (interface{}, error) {
	var (
	//tags, cacheTags []models.Tag
	)
	cache := cache_service.Tag{
		State:    t.State,
		PageNum:  t.PageNum,
		PageSize: t.PageSize,
	}
	_ = cache.GetTagsKey()
	//redis判断后续操作
	tags, err := models.GetTags(t.PageNum, t.PageSize, t.getMaps())
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (t *Tag) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	if t.Name != "" {
		maps["name"] = t.Name
	}
	if t.State >= 0 {
		maps["state"] = t.State
	}
	return maps
}
