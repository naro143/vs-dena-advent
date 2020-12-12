package handler

import (
	"net/http"

	"golang.org/x/sync/errgroup"

	"github.com/naro143/vs-dena-advent/model"
	"github.com/naro143/vs-dena-advent/qiita"

	"github.com/naro143/vs-dena-advent/model/repository"
)

type Handler struct {
	likesRepo   repository.Likes
	articleRepo repository.Article
}

func New(lr repository.Likes, as repository.Article) *Handler {
	return &Handler{
		likesRepo:   lr,
		articleRepo: as,
	}
}

func (h *Handler) GetLikes(w http.ResponseWriter, r *http.Request) {
	likes, err := h.likesRepo.GetNew()
	if err != nil {
		respondError(w, r, err, http.StatusInternalServerError, nil)
		return
	}
	respondSuccess(w, r, http.StatusOK, likes)
}

const (
	naiteiTitle    = "dena-21-shinsotsu"
	shinsotsuTitle = "dena-20-shinsotsu"
	generalTitle   = "dena"
)

func (h *Handler) UpdateLikes(w http.ResponseWriter, r *http.Request) {
	var likes model.Likes
	eg := &errgroup.Group{}
	eg.Go(func() error {
		naitei, err := qiita.GetAllLikes(2020, naiteiTitle)
		if err != nil {
			return err
		}
		likes.Naitei = naitei
		return nil
	})
	eg.Go(func() error {
		shinsotsu, err := qiita.GetAllLikes(2020, shinsotsuTitle)
		if err != nil {
			return err
		}
		likes.Shinsotsu = shinsotsu
		return nil
	})
	eg.Go(func() error {
		general, err := qiita.GetAllLikes(2020, generalTitle)
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

	if err := h.likesRepo.Create(&likes); err != nil {
		respondError(w, r, err, http.StatusInternalServerError, nil)
		return
	}
	respondSuccess(w, r, http.StatusCreated, nil)
}

func (h *Handler) ListArticles(w http.ResponseWriter, r *http.Request) {
	as, err := h.articleRepo.List()
	if err != nil {
		respondError(w, r, err, http.StatusInternalServerError, nil)
		return
	}
	respondSuccess(w, r, http.StatusOK, as)
}

func (h *Handler) UpdateArticles(w http.ResponseWriter, r *http.Request) {
	as := &model.Articles{}
	eg := &errgroup.Group{}
	eg.Go(func() error {
		naitei, err := qiita.GetArticles(2020, naiteiTitle)
		if err != nil {
			return err
		}
		as.Naitei = naitei
		return nil
	})
	eg.Go(func() error {
		shinsotsu, err := qiita.GetArticles(2020, shinsotsuTitle)
		if err != nil {
			return err
		}
		as.Shinsotsu = shinsotsu
		return nil
	})
	eg.Go(func() error {
		general, err := qiita.GetArticles(2020, generalTitle)
		if err != nil {
			return err
		}
		as.General = general
		return nil
	})
	if err := eg.Wait(); err != nil {
		respondError(w, r, err, http.StatusInternalServerError, nil)
		return
	}
	var naiteiTotalLikes int64
	for _, a := range as.Naitei {
		naiteiTotalLikes += a.Likes
	}
	as.NaiteiTotalLikes = naiteiTotalLikes
	var shinsotsuTotalLikes int64
	for _, a := range as.Shinsotsu {
		shinsotsuTotalLikes += a.Likes
	}
	as.ShinsotsuTotalLikes = shinsotsuTotalLikes
	var generalTotalLikes int64
	for _, a := range as.General {
		generalTotalLikes += a.Likes
	}
	as.GeneralTotalLikes = generalTotalLikes

	if err := h.articleRepo.Update(as); err != nil {
		respondError(w, r, err, http.StatusInternalServerError, nil)
		return
	}
	respondSuccess(w, r, http.StatusCreated, nil)
}
