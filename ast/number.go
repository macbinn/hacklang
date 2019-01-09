package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/builtin"
	"github.com/macbinn/hacklang/value"
)

type NumberNode struct {
	Value int
}

func (n *NumberNode) String() string {
	return fmt.Sprintf("<Number %d>", n.Value)
}

func (n *NumberNode) Eval(scope *value.Scope) value.Object {
	return builtin.NewNumber(n.Value)
}
