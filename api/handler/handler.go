package handler

import (
	"net/http"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"

	"github.com/tockn/vs-dena-advent/model"
	"github.com/tockn/vs-dena-advent/qiita"

	"github.com/tockn/vs-dena-advent/model/repository"
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
	shinsotsuTitle = "dena-20-shinsostu"
	generalTitle   = "dena"
)

func (h *Handler) UpdateLikes(w http.ResponseWriter, r *http.Request) {
	var likes model.Likes
	eg := &errgroup.Group{}
	eg.Go(func() error {
		shinsotsu, err := qiita.GetAllLikes(2019, shinsotsuTitle)
		if err != nil {
			return err
		}
		likes.Shinsotsu = shinsotsu
		return nil
	})
	eg.Go(func() error {
		general, err := qiita.GetAllLikes(2019, generalTitle)
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

	as, err := h.articleRepo.List()
	if err != nil {
		respondError(w, r, err, http.StatusInternalServerError, nil)
		return
	}
	wg := &sync.WaitGroup{}
	for i, a := range as.Shinsotsu {
		wg.Add(1)
		go func() {
			split := strings.Split(a.URL, "/")
			articleID := split[len(split)-1]
			likes, _ := qiita.GetLikesByArticleID(articleID)
			as.Shinsotsu[i].Likes = likes
			wg.Done()
		}()
	}
	for i, a := range as.General {
		wg.Add(1)
		go func() {
			split := strings.Split(a.URL, "/")
			articleID := split[len(split)-1]
			likes, _ := qiita.GetLikesByArticleID(articleID)
			as.General[i].Likes = likes
			wg.Done()
		}()
	}
	wg.Wait()
	if err := h.articleRepo.Update(as); err != nil {
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
		shinsotsu, err := qiita.GetArticles(2019, shinsotsuTitle)
		if err != nil {
			return err
		}
		as.Shinsotsu = shinsotsu
		return nil
	})
	eg.Go(func() error {
		general, err := qiita.GetArticles(2019, generalTitle)
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
	if err := h.articleRepo.Update(as); err != nil {
		respondError(w, r, err, http.StatusInternalServerError, nil)
		return
	}
	respondSuccess(w, r, http.StatusCreated, nil)
}
