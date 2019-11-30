package datastore

import (
	"errors"
	"time"

	"github.com/rs/xid"
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
	var ls []*model.Likes
	if _, err := r.client.GetAll(q, &ls); err != nil {
		return nil, err
	}
	if len(ls) == 0 {
		return nil, errors.New("not found")
	}
	return ls[0], nil
}

func (r *likesRepository) Create(likes *model.Likes) error {
	likes.ID = xid.New().String()
	likes.UpdatedAt = time.Now()
	if _, err := r.client.Put(likes); err != nil {
		return err
	}
	return nil
}
