package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/buildin"
)

type ExprList struct {
	Nodes []Node
}

func (e *ExprList) String() string {
	return fmt.Sprintf("<ExprList Nodes=%s>", e.Nodes)
}

func (e *ExprList) Eval(scope *Scope) buildin.Object {
	var obj buildin.Object
	for _, node := range e.Nodes {
		obj = node.Eval(scope)
	}
	return obj
}
