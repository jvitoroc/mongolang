package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Update interface {
	Filter(filter interface{}) Update
	Set(set bson.M) Update
	SetField(field string, value interface{}) Update
	Push(field string, value interface{}) Update
	Inc(field string, amount int64) Update

	GetCollection() *Collection
	One() (*mongo.UpdateResult, error)
	Many() (*mongo.UpdateResult, error)
}

type updateCommand struct {
	ctx  context.Context
	opts *options.UpdateOptions

	comm *Command

	filter bson.M
	update bson.M
}

func NewUpdateCommand() *updateCommand {
	return &updateCommand{filter: make(bson.M)}
}

func (u *updateCommand) Filter(filter interface{}) Update {
	u.comm.Filter(filter)
	return u
}

func (u *updateCommand) Set(set bson.M) Update {
	u.filter["$set"] = set
	return u
}

func (u *updateCommand) SetField(field string, value interface{}) Update {
	u.addUpdate(field, value, "$set")

	return u
}

func (u *updateCommand) Push(field string, value interface{}) Update {
	u.addUpdate(field, value, "$push")

	return u
}

func (u *updateCommand) Inc(field string, amount int64) Update {
	u.addUpdate(field, amount, "$ic")

	return u
}

func (u *updateCommand) addUpdate(key, value interface{}, field string) {
	v, ok := u.update[field]
	var target bson.M = v.(bson.M)
	if !ok {
		target = make(bson.M)
		u.filter[field] = target
	}

	target[field] = value
}

func (u *updateCommand) GetCollection() *Collection {
	return u.comm.coll
}

func (u *updateCommand) One() (*mongo.UpdateResult, error) {
	return u.GetCollection().collection.UpdateOne(u.getParameters())
}

func (u *updateCommand) Many() (*mongo.UpdateResult, error) {
	return u.GetCollection().collection.UpdateMany(u.getParameters())
}

func (u *updateCommand) getParameters() (context.Context, interface{}, bson.M, *options.UpdateOptions) {
	return u.ctx, u.filter, u.update, u.opts
}
