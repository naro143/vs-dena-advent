package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tockn/vs-dena-advent/api/server"
)

func New(s *server.Server) http.Handler {
	r := chi.NewRouter()
	r.Get("/likes", s.GetLikes)
	r.Post("/likes", s.UpdateLikes)
	return r
}
