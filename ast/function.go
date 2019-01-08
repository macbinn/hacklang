package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/builtin"
)

type FunctionNode struct {
	Arguments []string
	Body      []Node
	Scope     *Scope
}

func (f *FunctionNode) String() string {
	return fmt.Sprintf("<Function Arguments=%v, Body=%v>", f.Arguments, f.Body)
}

func (f *FunctionNode) Eval(scope *Scope) builtin.Object {
	f.Scope = NewScope(scope)
	return builtin.NewFunction("f", func(args ...builtin.Object) builtin.Object {
		for i, arg := range args {
			f.Scope.Register(f.Arguments[i], arg)
		}
		for _, node := range f.Body {
			node.Eval(f.Scope)
		}
		return nil
	})
}
