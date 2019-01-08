package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/buildin"
)

type FunctionNode struct {
	Arguments []string
	Body      []Node
	Scope     *Scope
}

func (f *FunctionNode) String() string {
	return fmt.Sprintf("<Function Arguments=%v, Body=%v>", f.Arguments, f.Body)
}

func (f *FunctionNode) Eval(scope *Scope) buildin.Object {
	f.Scope = NewScope(scope)
	return buildin.NewFunction("f", func(args ...buildin.Object) buildin.Object {
		for i, arg := range args {
			f.Scope.Register(f.Arguments[i], arg)
		}
		for _, node := range f.Body {
			node.Eval(f.Scope)
		}
		return nil
	})
}
