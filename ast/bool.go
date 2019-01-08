package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/buildin"
)

type BoolNode struct {
	Value bool
}

func (b *BoolNode) String() string {
	return fmt.Sprintf("<Bool %v>", b.Value)
}

func (b *BoolNode) Eval(scope *Scope) buildin.Object {
	return buildin.NewBool(b.Value)
}
