package dao

import (
	"go_api_framework/internal/model"
	"go_api_framework/pkg/app"
)

func (d *Dao) GetCategoryById(id uint32, state uint8) (model.Category, error) {
	category := model.Category{Model: &model.Model{Id: id}, State: state}
	return category.Get(d.engine)
}

func (d *Dao) GetCategoryList(page, pageSize int) ([]*model.Category, error) {
	category := model.Category{State: model.STATE_OPEN}
	pageOffset := app.GetPageOffset(page, pageSize)
	return category.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) GetCategoryListByName(name string, state uint8, page, pageSize int) ([]*model.Category, error) {
	category := model.Category{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return category.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) GetCategoryListByIds(ids []uint32, state uint8) ([]*model.Category, error) {
	category := model.Category{State: state}
	return category.ListByIds(d.engine, ids)
}

func (d *Dao) CountCategory(name string, state uint8) (int, error) {
	category := model.Category{Name: name, State: state}
	return category.Count(d.engine)
}

func (d *Dao) CreateCategory(name string, state uint8, createdBy string) error {
	category := model.Category{
		Name:  name,
		State: state,
		Model: &model.Model{
			CreatedBy: createdBy,
		},
	}

	return category.Create(d.engine)
}

func (d *Dao) UpdateCategory(id uint32, name string, state uint8, modifiedBy string) error {
	category := model.Category{
		Model: &model.Model{
			Id: id,
		},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}

	return category.Update(d.engine, values)
}

func (d *Dao) DeleteCategory(id uint32) error {
	category := model.Category{Model: &model.Model{Id: id}}
	return category.Delete(d.engine)
}

func (d *Dao) SearchCategory(key string) ([]*model.Category, error) {
	category := model.Category{
		Name:        key,
		Description: key,
	}
	return category.Search(d.engine)
}
