package main

type Command struct {
	coll *Collection

	filter interface{}
}

func (c *Command) Filter(filter interface{}) {
	c.filter = filter
}
