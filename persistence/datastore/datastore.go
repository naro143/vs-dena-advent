package datastore

import (
	"context"

	"go.mercari.io/datastore/clouddatastore"

	"go.mercari.io/datastore/boom"
)

type Client struct {
	*boom.Boom
}

func NewClient(ctx context.Context) (*Client, error) {
	ds, err := clouddatastore.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	return &Client{
		Boom: boom.FromClient(ctx, ds),
	}, nil
}
