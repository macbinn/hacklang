package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/value"
)

type ReturnNode struct {
	Expr Node
}

func (r *ReturnNode) String() string {
	return fmt.Sprintf("<Return Expr=%s>", r.Expr)
}

func (r *ReturnNode) Eval(scope *value.Scope) value.Object {
	return r.Expr.Eval(scope)
}

func (r *ReturnNode) Code() string {
	return fmt.Sprintf("return %s", r.Expr.Code())
}
