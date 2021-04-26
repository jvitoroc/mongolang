package expr

import "go.mongodb.org/mongo-driver/bson"

type ExprTree struct {
	Root Node
}

type Node interface {
	Build([]interface{}) interface{}
}

type And struct {
	Operands []Node
}

type Or struct {
	Operands []Node
}

type Gt struct {
	Field string
	Value Node
}

type Lt struct {
	Field string
	Value Node
}

type Gte struct {
	Field string
	Value Node
}

type Lte struct {
	Field string
	Value Node
}

type Literal struct {
	Value interface{}
}

type Variable struct {
	Index int
}

func NewExprTree(tokens []*Token) *ExprTree {
	e := &ExprTree{}

	return e
}

func (e *ExprTree) Build(vars []interface{}) interface{} {
	return e.Root.Build(vars)
}

func (n *And) Build(vars []interface{}) interface{} {
	built := []interface{}{}
	for _, oper := range n.Operands {
		built = append(built, oper.Build(vars))
	}

	return bson.M{
		"$and": built,
	}
}

func (n *Or) Build(vars []interface{}) interface{} {
	built := []interface{}{}
	for _, oper := range n.Operands {
		built = append(built, oper.Build(vars))
	}

	return bson.M{
		"$or": built,
	}
}

func (n *Gt) Build(vars []interface{}) interface{} {
	return bson.M{
		n.Field: bson.M{
			"$gt": n.Value.Build(vars),
		},
	}
}

func (n *Lt) Build(vars []interface{}) interface{} {
	return bson.M{
		n.Field: bson.M{
			"$lt": n.Value.Build(vars),
		},
	}
}

func (n *Gte) Build(vars []interface{}) interface{} {
	return bson.M{
		n.Field: bson.M{
			"$gte": n.Value.Build(vars),
		},
	}
}

func (n *Lte) Build(vars []interface{}) interface{} {
	return bson.M{
		n.Field: bson.M{
			"$lte": n.Value.Build(vars),
		},
	}
}

func (n *Literal) Build(vars []interface{}) interface{} {
	return n.Value
}

func (n *Variable) Build(vars []interface{}) interface{} {
	return vars[n.Index]
}
