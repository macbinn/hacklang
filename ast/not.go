package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/builtin"
	"github.com/macbinn/hacklang/value"
)

type NotNode struct {
	Expr Node
}

func (n *NotNode) String() string {
	return fmt.Sprintf("<Not Expr=%s>", n.Expr)
}

func (n *NotNode) Eval(scope *value.Scope) value.Object {
	if boolObject(n.Expr.Eval(scope)) {
		return builtin.BoolFalse
	}
	return builtin.BoolTrue
}

func (n *NotNode) Code() string {
	return fmt.Sprintf("not %s", n.Expr.Code())
}
