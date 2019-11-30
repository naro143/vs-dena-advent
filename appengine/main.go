package main

import (
	"context"
	"net/http"

	"google.golang.org/appengine"

	"github.com/tockn/vs-dena-advent/api/router"
	"github.com/tockn/vs-dena-advent/api/server"

	"github.com/tockn/vs-dena-advent/persistence/datastore"
	"go.mercari.io/datastore/boom"
	"go.mercari.io/datastore/clouddatastore"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		panic(err)
	}
}

func run(ctx context.Context) error {
	ds, err := clouddatastore.FromContext(ctx)
	if err != nil {
		return err
	}
	c := boom.FromClient(ctx, ds)
	likeRepo := datastore.NewLikesRepository(c)
	s := server.NewServer(likeRepo)
	r := router.New(s)
	http.Handle("/", r)
	appengine.Main()
	return nil
}
