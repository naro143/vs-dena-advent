package datastore

import (
	"time"

	"github.com/tockn/vs-dena-advent/model/repository"

	"github.com/tockn/vs-dena-advent/model"
	"go.mercari.io/datastore/boom"
)

type likesRepository struct {
	client *boom.Boom
}

func NewLikesRepository(c *boom.Boom) repository.Likes {
	return &likesRepository{
		client: c,
	}
}

func (r *likesRepository) GetNew() (*model.Likes, error) {
	q := r.client.NewQuery("Likes").
		Order("UpdatedAt").
		Limit(1)
	ls := make([]*model.Likes, 1)
	if _, err := r.client.GetAll(q, ls); err != nil {
		return nil, err
	}
	return ls[0], nil
}

func (r *likesRepository) Update(likes *model.Likes) error {
	likes.UpdatedAt = time.Now()
	if _, err := r.client.Put(likes); err != nil {
		return err
	}
	return nil
}
