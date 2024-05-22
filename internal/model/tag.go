package model

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

// gorm的所需方法
func (t Tag) TableName() string {
	return "blog_tag"
}
