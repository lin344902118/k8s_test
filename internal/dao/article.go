package dao

import (
	"go_api_framework/internal/model"
)

type Article struct {
	Id            uint32 `json:"id"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         uint8  `json:"state"`
}

func (d *Dao) CreateArticle(param *Article) (*model.Article, error) {
	article := model.Article{
		Title:         param.Title,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		Model:         &model.Model{CreatedBy: param.CreatedBy},
	}
	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(param *Article) error {
	article := model.Article{Model: &model.Model{Id: param.Id}}
	values := map[string]interface{}{
		"modified_by": param.ModifiedBy,
		"state":       param.State,
	}
	if param.Title != "" {
		values["title"] = param.Title
	}
	if param.CoverImageUrl != "" {
		values["cover_image_url"] = param.CoverImageUrl
	}
	if param.Content != "" {
		values["content"] = param.Content
	}

	return article.Update(d.engine, values)
}

func (d *Dao) GetArticle(id uint32, state uint8) (model.Article, error) {
	article := model.Article{Model: &model.Model{Id: id}, State: state}
	return article.Get(d.engine)
}

func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{Model: &model.Model{Id: id}}
	return article.Delete(d.engine)
}

func (d *Dao) GetArticleList(page, pageSize int) ([]*model.Article, error) {
	article := model.Article{State: model.STATE_OPEN}
	var pageOffset = 0
	if page > 0 {
		pageOffset = (page - 1) * pageSize
	}
	return article.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) SearchArticle(key string) ([]*model.Article, error) {
	article := model.Article{Title: key, Content: key}
	return article.Search(d.engine)
}
