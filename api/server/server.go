package server

import (
	"net/http"

	"golang.org/x/sync/errgroup"

	"github.com/tockn/vs-dena-advent/model"
	"github.com/tockn/vs-dena-advent/qiita"

	"github.com/tockn/vs-dena-advent/model/repository"
)

type Server struct {
	likesRepo repository.Likes
}

func NewServer(lr repository.Likes) *Server {
	return &Server{
		likesRepo: lr,
	}
}

func (s *Server) GetLikes(w http.ResponseWriter, r *http.Request) {
	likes, err := s.likesRepo.GetNew()
	if err != nil {
		respondError(w, r, err, http.StatusInternalServerError, nil)
		return
	}
	respondSuccess(w, r, http.StatusOK, likes)
}

const (
	shinsotsuTitle = "dena-20-shinsostu"
	generalTitle   = "dena"
)

func (s *Server) UpdateLikes(w http.ResponseWriter, r *http.Request) {
	var likes model.Likes
	eg := &errgroup.Group{}
	eg.Go(func() error {
		shinsotsu, err := qiita.GetLikes(2019, shinsotsuTitle)
		if err != nil {
			return err
		}
		likes.Shinsotsu = shinsotsu
		return nil
	})
	eg.Go(func() error {
		general, err := qiita.GetLikes(2019, generalTitle)
		if err != nil {
			return err
		}
		likes.General = general
		return nil
	})
	if err := eg.Wait(); err != nil {
		respondError(w, r, err, http.StatusInternalServerError, nil)
		return
	}

	if err := s.likesRepo.Create(&likes); err != nil {
		respondError(w, r, err, http.StatusInternalServerError, nil)
		return
	}
	respondSuccess(w, r, http.StatusCreated, nil)
}
