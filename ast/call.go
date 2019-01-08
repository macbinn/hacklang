package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/buildin"
)

type CallNode struct {
	Callee    Node
	Arguments []Node
}

func (c *CallNode) String() string {
	return fmt.Sprintf("<Call Callee=%s Arguments=%s>", c.Callee, c.Arguments)
}

func (c *CallNode) Eval(scope *Scope) buildin.Object {
	var args []buildin.Object
	for _, arg := range c.Arguments {
		args = append(args, arg.Eval(scope))
	}
	calleeFunc := c.Callee.Eval(scope)
	fn := calleeFunc.(*buildin.Function)
	return fn.Func()(args...)
}
