package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/builtin"
	"github.com/macbinn/hacklang/value"
)

type StringNode struct {
	Value string
}

func (s *StringNode) String() string {
	return fmt.Sprintf("<String %s>", s.Value)
}

func (s *StringNode) Eval(scope *value.Scope) value.Object {
	return builtin.NewString(s.Value)
}
