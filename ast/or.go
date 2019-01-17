package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/builtin"
	"github.com/macbinn/hacklang/value"
)

type OrNode struct {
	Left Node
	Right Node
}

func (o *OrNode) String() string {
	return fmt.Sprintf("<Or Left=%s Right=%s>", o.Left, o.Right)
}

func (o *OrNode) Eval(scope *value.Scope) value.Object {
	if boolObject(o.Left.Eval(scope)) {
		return builtin.BoolTrue
	}
	if boolObject(o.Right.Eval(scope)) {
		return builtin.BoolTrue
	}
	return builtin.BoolFalse
}

func (o *OrNode) Code() string {
	return fmt.Sprintf("%s or %s", o.Left.Code(), o.Right.Code())
}
