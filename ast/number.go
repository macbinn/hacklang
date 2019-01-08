package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/builtin"
)

type NumberNode struct {
	Value int
}

func (n *NumberNode) String() string {
	return fmt.Sprintf("<Number %d>", n.Value)
}

func (n *NumberNode) Eval(scope *Scope) builtin.Object {
	return builtin.NewNumber(n.Value)
}
