package main

import (
	"log"
	"net/http"

	"github.com/naro143/vs-dena-advent/api/handler"
	"github.com/naro143/vs-dena-advent/api/router"
	"github.com/naro143/vs-dena-advent/persistence/memory"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	likeRepo := memory.NewLikesRepository()
	articleRepo := memory.NewArticleRepository()
	s := handler.New(likeRepo, articleRepo)
	r := router.New(s)
	log.Println("serving...")
	return http.ListenAndServe(":8080", r)
}
