package tag_service

import "example.com/m/v2/models"

type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

func (t *Tag) Add() (interface{}, error) {
	return models.AddTag(t.Name, t.State, t.CreatedBy)
}
