package service

import "go_api_framework/pkg/convert"

type Category struct {
	Id          uint32 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	State       uint8  `json:"state"`
	PublishOn   string `json:"publishOn"`
}

func (svc *Service) GetCategoryList(page, pageSize int) ([]*Category, error) {
	categorys, err := svc.dao.GetCategoryList(page, pageSize)
	if err != nil {
		return nil, err
	}
	datas := make([]*Category, 0, len(categorys))
	for _, category := range categorys {
		c := &Category{
			Id:          category.Id,
			Name:        category.Name,
			Description: category.Description,
			State:       category.State,
			PublishOn:   convert.UnixTime(category.CreatedOn).Date(),
		}
		datas = append(datas, c)
	}
	return datas, nil
}
