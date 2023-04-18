package dao

import "go_api_framework/internal/model"

func (d *Dao) GetArticleCategoryByAId(articleId uint32) (model.ArticleCategory, error) {
	articleCategory := model.ArticleCategory{ArticleId: articleId}
	return articleCategory.GetByAId(d.engine)
}

func (d *Dao) GetArticleCategoryListByAId(articleId uint32) ([]*model.ArticleCategory, error) {
	articleCategory := model.ArticleCategory{ArticleId: articleId}
	return articleCategory.ListByAId(d.engine)
}

func (d *Dao) GetArticleCategoryListByTId(categoryId uint32) ([]*model.ArticleCategory, error) {
	articleCategory := model.ArticleCategory{CategoryId: categoryId}
	return articleCategory.ListByTId(d.engine)
}

func (d *Dao) GetArticleCategoryListByAIds(articleIds []uint32) ([]*model.ArticleCategory, error) {
	articleCategory := model.ArticleCategory{}
	return articleCategory.ListByAIds(d.engine, articleIds)
}

func (d *Dao) CreateArticleCategory(articleId, categoryId uint32, createdBy string) error {
	articleCategory := model.ArticleCategory{
		Model: &model.Model{
			CreatedBy: createdBy,
		},
		ArticleId:  articleId,
		CategoryId: categoryId,
	}
	return articleCategory.Create(d.engine)
}

func (d *Dao) UpdateArticleCategory(articleId, categoryId uint32, modifiedBy string) error {
	articleCategory := model.ArticleCategory{ArticleId: articleId}
	values := map[string]interface{}{
		"article_id":  articleId,
		"category_id": categoryId,
		"modified_by": modifiedBy,
	}
	return articleCategory.UpdateOne(d.engine, values)
}

func (d *Dao) DeleteArticleCategory(articleId uint32) error {
	articleCategory := model.ArticleCategory{ArticleId: articleId}
	return articleCategory.DeleteOne(d.engine)
}
