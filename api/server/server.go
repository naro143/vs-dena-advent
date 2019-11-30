package server

import (
	"net/http"

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
		respondError(w, err, http.StatusInternalServerError, nil)
		return
	}
	respondSuccess(w, http.StatusOK, likes)
}

const (
	shinsotsuTitle = "dena-20-shinsostu"
	generalTitle   = "dena"
)

func (s *Server) UpdateLikes(w http.ResponseWriter, r *http.Request) {
	shinsotsu, err := qiita.GetLikes(2019, shinsotsuTitle)
	if err != nil {
		respondError(w, err, http.StatusInternalServerError, nil)
		return
	}
	general, err := qiita.GetLikes(2019, generalTitle)
	if err != nil {
		respondError(w, err, http.StatusInternalServerError, nil)
		return
	}
	likes := &model.Likes{
		Shinsotsu: shinsotsu,
		General:   general,
	}
	if err := s.likesRepo.Update(likes); err != nil {
		respondError(w, err, http.StatusInternalServerError, nil)
		return
	}
	respondSuccess(w, http.StatusCreated, nil)
}
