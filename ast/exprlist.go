package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/value"
)

type ExprList struct {
	Nodes []Node
}

func (e *ExprList) String() string {
	return fmt.Sprintf("<ExprList Nodes=%s>", e.Nodes)
}

func (e *ExprList) Eval(scope *value.Scope) value.Object {
	var obj value.Object
	for _, node := range e.Nodes {
		obj = node.Eval(scope)
	}
	return obj
}
