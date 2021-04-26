package main

type Command struct {
	coll *Collection

	filter interface{}
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Filter(filter interface{}) {
	c.filter = filter
}
