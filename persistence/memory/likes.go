package memory

import (
	"time"

	"github.com/tockn/vs-dena-advent/model"
	"github.com/tockn/vs-dena-advent/model/repository"
)

type likesRepository struct {
	memory []*model.Likes
}

func NewLikesRepository() repository.Likes {
	mem := make([]*model.Likes, 1)
	mem[0] = &model.Likes{
		Shinsotsu: 100,
		General:   101,
		UpdatedAt: time.Now(),
	}
	return &likesRepository{
		memory: mem,
	}
}

func (r *likesRepository) GetNew() (*model.Likes, error) {
	return r.memory[len(r.memory)-1], nil
}

func (r *likesRepository) Update(likes *model.Likes) error {
	likes.UpdatedAt = time.Now()
	r.memory = append(r.memory, likes)
	return nil
}
