package repository

import "github.com/tockn/vs-dena-advent/model"

type Likes interface {
	GetNew() (*model.Likes, error)
	Update(likes *model.Likes) error
}
