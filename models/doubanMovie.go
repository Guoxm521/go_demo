package models

type DoubanMovie struct {
	Model
	Title    string `json:"title",des:"电影名称"`
	Time     string `json:"time'",des:"上映时间"`
	Duration string `json:"duration",des:"播放时长"`
	Director string `json:"director",des:"导演"`
	FilmType string `json:"filmType",des:"类型"`
	Address  string `json:"address",des:"制片国家/地区"`
	Language string `json:"language",des:"语言"`
	Des      string `json:"des",des:"描述"`
}

func (_this *DoubanMovie) AddMovie() (interface{}, error) {
	result := db.Create(&_this)
	if result.Error != nil {
		return nil, result.Error
	}
	return "success1232131", nil
}

func (_this *DoubanMovie) List() (interface{}, error) {
	var (
		movie []DoubanMovie
		err   error
	)
	err = db.Offset(0).Limit(1000).Find(&movie).Error
	return movie, err
}

func (_this *DoubanMovie) add() {

}
