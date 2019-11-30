package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tockn/vs-dena-advent/api/handler"
)

func New(h *handler.Handler) http.Handler {
	r := chi.NewRouter()
	r.Use(cors)
	r.Get("/likes", h.GetLikes)
	r.Get("/articles", h.ListArticles)

	// for cron job
	r.Get("/updatelikes", h.UpdateLikes)
	r.Get("/updatearticles", h.UpdateArticles)
	return r
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
