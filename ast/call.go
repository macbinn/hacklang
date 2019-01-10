package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/builtin"
	"github.com/macbinn/hacklang/value"
	"strings"
)

type CallNode struct {
	Callee    Node
	Arguments []Node
}

func (c *CallNode) Code() string {
	var args []string
	for _, node := range c.Arguments {
		args = append(args, node.Code())
	}
	argsCode := strings.Join(args, ", ")
	return fmt.Sprintf("%s(%s)", c.Callee.Code(), argsCode)
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
