package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/buildin"
)

type NumberNode struct {
	Value int
}

func (n *NumberNode) String() string {
	return fmt.Sprintf("<Number %d>", n.Value)
}

func (n *NumberNode) Eval(scope *Scope) buildin.Object {
	return buildin.NewNumber(n.Value)
}
