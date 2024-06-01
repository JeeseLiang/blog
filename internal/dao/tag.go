package dao

import (
	"blog/internal/model"
	"blog/pkg/app"
	"time"
)

func (d *Dao) CountTag(name string, state uint8) (int64, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) CreateTag(name string, state uint8, createBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{
			ModifiedOn: uint32(time.Now().Unix()),
			CreatedOn:  uint32(time.Now().Unix()),
			CreatedBy:  createBy,
		},
	}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{
			ModifiedBy: modifiedBy,
			ModifiedOn: uint32(time.Now().Unix()),
			ID:         id,
		},
	}
	return tag.Update(d.engine)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{
		Model: &model.Model{
			ID: id,
		},
	}
	return tag.Delete(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	return model.Tag{
		Name:  name,
		State: state,
	}.List(d.engine, app.GetPageOffset(page, pageSize), pageSize)
}
