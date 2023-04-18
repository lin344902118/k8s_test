package service

import (
	"go_api_framework/internal/dao"
	"go_api_framework/internal/model"
	"go_api_framework/pkg/convert"
	"go_api_framework/pkg/errcode"
)

type Article struct {
	Id            uint32            `json:"id"`
	Title         string            `json:"title"`
	Content       string            `json:"content"`
	CoverImageUrl string            `json:"cover_image_url"`
	State         uint8             `json:"state"`
	PublishOn     string            `json:"publishOn"`
	Categorys     []*model.Category `json:"category"`
}

func (svc *Service) GetAllArticleWithCategorys(page, pageSize int) ([]*Article, error) {
	articles, err := svc.dao.GetArticleList(page, pageSize)
	results := make([]*Article, 0, len(articles))
	if err != nil {
		return nil, err
	}
	for _, article := range articles {
		acs, err := svc.dao.GetArticleCategoryListByAId(article.Id)
		if err != nil {
			return nil, err
		}
		categoryIds := make([]uint32, len(acs))
		for _, ac := range acs {
			categoryIds = append(categoryIds, ac.CategoryId)
		}
		categorys, err := svc.dao.GetCategoryListByIds(categoryIds, 1)
		if err != nil {
			return nil, err
		}
		a := &Article{
			Id:            article.Id,
			Title:         article.Title,
			Content:       article.Content,
			CoverImageUrl: article.CoverImageUrl,
			State:         article.State,
			PublishOn:     convert.UnixTime(article.CreatedOn).Date(),
			Categorys:     categorys,
		}
		results = append(results, a)
	}
	return results, nil
}

func (svc *Service) GetArticleList(page, pageSize int) ([]*Article, error) {
	articles, err := svc.dao.GetArticleList(page, pageSize)
	if err != nil {
		return nil, err
	}
	results := make([]*Article, 0, len(articles))
	for _, article := range articles {
		a := &Article{
			Id:            article.Id,
			Title:         article.Title,
			Content:       article.Content,
			CoverImageUrl: article.CoverImageUrl,
			State:         article.State,
			PublishOn:     convert.UnixTime(article.CreatedOn).Date(),
		}
		results = append(results, a)
	}
	return results, nil
}

func (svc *Service) Search(key string) ([]*Article, error) {
	articles, err := svc.dao.SearchArticle(key)
	if err != nil {
		return nil, err
	}
	results := make([]*Article, 0, len(articles))
	for _, article := range articles {
		acs, err := svc.dao.GetArticleCategoryListByAId(article.Id)
		if err != nil {
			return nil, err
		}
		categoryIds := make([]uint32, len(acs))
		for _, ac := range acs {
			categoryIds = append(categoryIds, ac.CategoryId)
		}
		categorys, err := svc.dao.GetCategoryListByIds(categoryIds, 1)
		if err != nil {
			return nil, err
		}
		a := &Article{
			Id:            article.Id,
			Title:         article.Title,
			Content:       article.Content,
			CoverImageUrl: article.CoverImageUrl,
			State:         article.State,
			PublishOn:     convert.UnixTime(article.CreatedOn).Date(),
			Categorys:     categorys,
		}
		results = append(results, a)
	}
	return results, nil
}

func (svc *Service) GetArticle(id uint32) (*Article, error) {
	article, err := svc.dao.GetArticle(id, model.STATE_OPEN)
	if err != nil {
		return nil, err
	}
	acs, err := svc.dao.GetArticleCategoryListByAId(article.Id)
	if err != nil {
		return nil, err
	}
	categoryIds := make([]uint32, len(acs))
	for _, ac := range acs {
		categoryIds = append(categoryIds, ac.CategoryId)
	}
	categorys, err := svc.dao.GetCategoryListByIds(categoryIds, 1)
	if err != nil {
		return nil, err
	}
	a := &Article{
		Id:            article.Id,
		Title:         article.Title,
		Content:       article.Content,
		CoverImageUrl: article.CoverImageUrl,
		State:         article.State,
		PublishOn:     convert.UnixTime(article.CreatedOn).Date(),
		Categorys:     categorys,
	}
	return a, nil
}

func (svc *Service) UpdateArticle(article *Article) error {
	if article.Id == 0 {
		return errcode.MissingArticleId
	}
	param := &dao.Article{
		Id: article.Id,
	}
	if article.Title != "" {
		param.Title = article.Title
	}
	if article.Content != "" {
		param.Content = article.Content
	}
	if article.CoverImageUrl != "" {
		param.CoverImageUrl = article.CoverImageUrl
	}
	return svc.dao.UpdateArticle(param)
}

func (svc *Service) CreateArticle(article Article) (uint32, error) {
	if article.Title == "" {
		return 0, errcode.ArticleTitleRequired
	}
	if article.Content == "" {
		return 0, errcode.ArticleContentRequired
	}
	param := &dao.Article{
		Title:         article.Title,
		Content:       article.Content,
		CoverImageUrl: article.CoverImageUrl,
		State:         model.STATE_OPEN,
	}
	a, err := svc.dao.CreateArticle(param)
	return a.Id, err
}

func (svc *Service) DeleteArticle(id uint32) error {
	return svc.dao.DeleteArticle(id)
}
