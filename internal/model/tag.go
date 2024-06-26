package model

import "gorm.io/gorm"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

// gorm的所需方法
func (t Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Count(db *gorm.DB) (int64, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	err := db.Model(&t).Where("is_del - ?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, err
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? and is_del = ?", t.Model.ID, 0).Delete(&t).Error
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB) error {
	return db.Model(&Tag{}).Where("id = ? and is_del = ?", t.ID, 0).Updates(t).Error
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	tags := []*Tag{}
	if pageOffset > 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ? and is_del = ?", t.State, 0)
	err := db.Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}
