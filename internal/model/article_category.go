package model

import "github.com/jinzhu/gorm"

type ArticleCategory struct {
	*Model
	ArticleId  uint32 `json:"article_id"`
	CategoryId uint32 `json:"category_id"`
}

func (a ArticleCategory) TableName() string {
	return "blog_article_category"
}

func (a ArticleCategory) GetByAId(db *gorm.DB) (ArticleCategory, error) {
	var articleCategory ArticleCategory
	err := db.Where("article_id = ? AND is_del = ?", a.ArticleId, 0).First(&articleCategory).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return articleCategory, err
	}

	return articleCategory, nil
}

func (a ArticleCategory) ListByAId(db *gorm.DB) ([]*ArticleCategory, error) {
	var articleCategorys []*ArticleCategory
	err := db.Where("article_id = ? AND is_del = ?", a.ArticleId, 0).Find(&articleCategorys).Error
	if err != nil {
		return nil, err
	}
	return articleCategorys, nil
}

func (a ArticleCategory) ListByTId(db *gorm.DB) ([]*ArticleCategory, error) {
	var articleCategorys []*ArticleCategory
	if err := db.Where("category_id = ? AND is_del = ?", a.ArticleId, 0).Find(&articleCategorys).Error; err != nil {
		return nil, err
	}

	return articleCategorys, nil
}

func (a ArticleCategory) ListByAIds(db *gorm.DB, articleIDs []uint32) ([]*ArticleCategory, error) {
	var articleCategorys []*ArticleCategory
	err := db.Where("article_id IN (?) AND is_del = ?", articleIDs, 0).Find(&articleCategorys).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articleCategorys, nil
}

func (a ArticleCategory) Create(db *gorm.DB) error {
	if err := db.Create(&a).Error; err != nil {
		return err
	}

	return nil
}

func (a ArticleCategory) UpdateOne(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Where("article_id = ? AND is_del = ?", a.ArticleId, 0).Limit(1).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func (a ArticleCategory) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", a.Model.Id, 0).Delete(&a).Error; err != nil {
		return err
	}

	return nil
}

func (a ArticleCategory) DeleteOne(db *gorm.DB) error {
	if err := db.Where("article_id = ? AND is_del = ?", a.ArticleId, 0).Delete(&a).Limit(1).Error; err != nil {
		return err
	}

	return nil
}
