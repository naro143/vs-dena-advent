package memory

import (
	"github.com/tockn/vs-dena-advent/model"
	"github.com/tockn/vs-dena-advent/model/repository"
)

type articleRepository struct {
}

func NewArticleRepository() repository.Article {
	return &articleRepository{}
}

func (r *articleRepository) List() (*model.Articles, error) {
	return &model.Articles{
		Shinsotsu: []model.Article{
			{
				URL: "https://qiita.com/Kuniwak/items/7e0e3f1cb6f3ae822215",
			},
		},
		General: []model.Article{
			{
				URL: "https://qiita.com/Kuniwak/items/7e0e3f1cb6f3ae822215",
			},
		},
	}, nil
}

func (r *articleRepository) Update(as *model.Articles) error {
	return nil
}
