package models

type Tag struct {
	Model
	Id         int
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}
type Data struct {
	Id         int    `json:"id"`
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
	//util.Console(result)
	back := Data{
		Id: result.Value.id,
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return back, nil
}
