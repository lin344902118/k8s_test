package model

import "github.com/jinzhu/gorm"

type Category struct {
	*Model
	Name        string `json:"name"`
	Description string `json:"description"`
	State       uint8  `json:"state"`
}

func (c Category) TableName() string {
	return "blog_category"
}

func (c Category) Count(db *gorm.DB) (int, error) {
	var count int
	if c.Name != "" {
		db = db.Where("name = ?", c.Name)
	}
	db = db.Where("state = ?", c.State)
	if err := db.Model(&c).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (c Category) List(db *gorm.DB, pageOffset, pageSize int) ([]*Category, error) {
	var categorys []*Category
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if c.Name != "" {
		db = db.Where("name = ?", c.Name)
	}
	db = db.Where("state = ?", c.State)
	if err = db.Where("is_del = ?", 0).Find(&categorys).Error; err != nil {
		return nil, err
	}

	return categorys, nil
}

func (c Category) ListByIds(db *gorm.DB, ids []uint32) ([]*Category, error) {
	var categorys []*Category
	db = db.Where("state = ? AND is_del = ?", c.State, 0)
	err := db.Where("id IN (?)", ids).Find(&categorys).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return categorys, nil
}

func (c Category) Get(db *gorm.DB) (Category, error) {
	var category Category
	err := db.Where("id = ? AND is_del = ? AND state = ?", c.Id, 0, c.State).First(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return category, err
	}

	return category, nil
}

func (c Category) Create(db *gorm.DB) error {
	return db.Create(&c).Error
}

func (c Category) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(c).Where("id = ? AND is_del = ?", c.Id, 0).Update(values).Error; err != nil {
		return err
	}
	return nil
}

func (c Category) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", c.Model.Id, 0).Delete(&c).Error
}

func (c Category) Search(db *gorm.DB) ([]*Category, error) {
	var categorys []*Category
	if c.Name != "" {
		db.Where("name like %?%", c.Name)
	}
	if c.Description != "" {
		db.Where("description like %?%", c.Description)
	}
	if err := db.Where("is_del = ?", 0).Find(&categorys).Error; err != nil {
		return nil, err
	}
	return categorys, nil
}
