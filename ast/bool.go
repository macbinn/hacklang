package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/builtin"
)

type BoolNode struct {
	Value bool
}

func (b *BoolNode) String() string {
	return fmt.Sprintf("<Bool %v>", b.Value)
}

func (b *BoolNode) Eval(scope *Scope) builtin.Object {
	return builtin.NewBool(b.Value)
}
