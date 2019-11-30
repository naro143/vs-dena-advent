package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tockn/vs-dena-advent/api/handler"
)

func New(h *handler.Handler) http.Handler {
	r := chi.NewRouter()
	r.Get("/likes", h.GetLikes)

	// for cron job
	r.Get("/updatelikes", h.UpdateLikes)
	return r
}
