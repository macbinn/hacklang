package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/buildin"
)

type StringNode struct {
	Value string
}

func (s *StringNode) String() string {
	return fmt.Sprintf("<String %s>", s.Value)
}

func (s *StringNode) Eval(scope *Scope) buildin.Object {
	return buildin.NewString(s.Value)
}
