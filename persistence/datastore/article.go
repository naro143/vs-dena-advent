package datastore

import (
	"github.com/naro143/vs-dena-advent/model"
	"github.com/naro143/vs-dena-advent/model/repository"
	"go.mercari.io/datastore/boom"
)

type articleRepository struct {
	client *boom.Boom
}

func NewArticleRepository(c *boom.Boom) repository.Article {
	return &articleRepository{
		client: c,
	}
}

func (r *articleRepository) List() (*model.Articles, error) {
	as := &model.Articles{}
	as.ID = 1
	if err := r.client.Get(as); err != nil {
		return nil, err
	}
	return as, nil
}

func (r *articleRepository) Update(as *model.Articles) error {
	as.ID = 1
	as.SetOpens()
	as.SetDays()
	if _, err := r.client.Put(as); err != nil {
		return err
	}
	return nil
}
