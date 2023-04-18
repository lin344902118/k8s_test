package model

import (
	"github.com/jinzhu/gorm"
	"go_api_framework/pkg/app"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	CoverImageUrl string `json:"cover_image_url"`
	Content       string `json:"content"`
	State         uint8  `json:"state"`
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Count(db *gorm.DB) (int, error) {
	var count int
	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}
	db = db.Where("state = ?", a.State)
	if err := db.Model(&a).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (a Article) Create(db *gorm.DB) (*Article, error) {
	if err := db.Create(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (a Article) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Where("id = ? AND is_del = ?", a.Id, 0).Update(values).Error; err != nil {
		return err
	}
	return nil
}

func (a Article) Get(db *gorm.DB) (Article, error) {
	var article Article
	db = db.Where("id = ? AND state = ? AND is_del = ?", a.Id, a.State, 0)
	err := db.First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return article, err
	}

	return article, nil
}

func (a Article) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", a.Id, 0).Delete(&a).Error; err != nil {
		return err
	}
	return nil
}

func (a Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var articles []*Article
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	db = db.Where("state = ?", a.State)
	if err = db.Where("is_del = ?", 0).Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

func (a Article) Search(db *gorm.DB) ([]*Article, error) {
	var articles []*Article
	if a.Title != "" {
		db = db.Where("title like ?", "%"+a.Title+"%")
	}
	if err := db.Where("is_del = ?", 0).Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}
