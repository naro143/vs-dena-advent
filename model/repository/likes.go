package repository

import "github.com/naro143/vs-dena-advent/model"

type Likes interface {
	GetNew() (*model.Likes, error)
	Create(likes *model.Likes) error
}

type Article interface {
	List() (*model.Articles, error)
	Update(as *model.Articles) error
}
