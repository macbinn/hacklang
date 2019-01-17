package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/builtin"
	"github.com/macbinn/hacklang/value"
	"strings"
)

type FunctionNode struct {
	Arguments []string
	Body      Node
}

func (f *FunctionNode) Code() string {
	args := strings.Join(f.Arguments, ", ")
	return fmt.Sprintf("(%s) => {\n%s\n}", args, f.Body.Code())
}

func (f *FunctionNode) String() string {
	return fmt.Sprintf("<Function Arguments=%v, Body=%v>", f.Arguments, f.Body)
}

func (f *FunctionNode) Eval(scope *value.Scope) value.Object {
	return builtin.NewFunction("", func(args ...value.Object) value.Object {
		fnScope := value.NewScope(scope)
		for i, arg := range args {
			if i < len(f.Arguments) {
				fnScope.Register(f.Arguments[i], arg)
			}
		}
		return f.Body.Eval(fnScope)
	})
}
