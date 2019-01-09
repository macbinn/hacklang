package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/builtin"
	"github.com/macbinn/hacklang/value"
)

type BoolNode struct {
	Value bool
}

func (b *BoolNode) String() string {
	return fmt.Sprintf("<Bool %v>", b.Value)
}

func (b *BoolNode) Eval(scope *value.Scope) value.Object {
	return builtin.NewBool(b.Value)
}
