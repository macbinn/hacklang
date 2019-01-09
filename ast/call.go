package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/builtin"
	"github.com/macbinn/hacklang/value"
)

type CallNode struct {
	Callee    Node
	Arguments []Node
}

func (c *CallNode) String() string {
	return fmt.Sprintf("<Call Callee=%s Arguments=%s>", c.Callee, c.Arguments)
}

func (c *CallNode) Eval(scope *value.Scope) value.Object {
	var args []value.Object
	for _, arg := range c.Arguments {
		args = append(args, arg.Eval(scope))
	}
	calleeFunc := c.Callee.Eval(scope)
	fn := calleeFunc.(*builtin.Function)
	return fn.Func()(args...)
}
