package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/builtin"
	"github.com/macbinn/hacklang/value"
	"strings"
)

type FunctionNode struct {
	Arguments []string
	Body      []Node
	Scope     *value.Scope
}

func (f *FunctionNode) Code() string {
	args := strings.Join(f.Arguments, ", ")
	var bodys []string
	for _, node := range f.Body {
		bodys = append(bodys, "  " + node.Code())
	}
	body := strings.Join(bodys, "\n")
	return fmt.Sprintf("(%s) => {\n%s\n}", args, body)
}

func (f *FunctionNode) String() string {
	return fmt.Sprintf("<Function Arguments=%v, Body=%v>", f.Arguments, f.Body)
}

func (f *FunctionNode) Eval(scope *value.Scope) value.Object {
	f.Scope = value.NewScope(scope)
	return builtin.NewFunction("f", func(args ...value.Object) value.Object {
		for i, arg := range args {
			if i < len(f.Arguments) {
				f.Scope.Register(f.Arguments[i], arg)
			}
		}
		for _, node := range f.Body {
			node.Eval(f.Scope)
		}
		return nil
	})
}
