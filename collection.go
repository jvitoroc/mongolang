package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct {
	collection *mongo.Collection
}

func (c *Collection) Update(ctx context.Context, opts *options.UpdateOptions) Update {
	return NewUpdateCommand()
}
