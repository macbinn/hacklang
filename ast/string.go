package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/builtin"
)

type StringNode struct {
	Value string
}

func (s *StringNode) String() string {
	return fmt.Sprintf("<String %s>", s.Value)
}

func (s *StringNode) Eval(scope *Scope) builtin.Object {
	return builtin.NewString(s.Value)
}
