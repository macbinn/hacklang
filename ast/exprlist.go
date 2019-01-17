package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/value"
	"strings"
)

type ExprList struct {
	Nodes []Node
}

func (e *ExprList) Code() string {
	var exprs []string
	for _, node := range e.Nodes {
		exprs = append(exprs, node.Code())
	}
	return strings.Join(exprs, "\n")
}

func (e *ExprList) String() string {
	return fmt.Sprintf("<ExprList Nodes=%s>", e.Nodes)
}

func (e *ExprList) Eval(scope *value.Scope) value.Object {
	var obj value.Object
	for _, node := range e.Nodes {
		obj = node.Eval(scope)
		if scope.Ret {
			return obj
		}
		if _, ok := node.(*ReturnNode); ok {
			scope.Ret = true
			return obj
		}
	}
	return obj
}
