package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/builtin"
	"github.com/macbinn/hacklang/value"
)

type AndNode struct {
	Left Node
	Right Node
}

func (a *AndNode) String() string {
	return fmt.Sprintf("<And Left=%s Right=%s>", a.Left, a.Right)
}

func (a *AndNode) Eval(scope *value.Scope) value.Object {
	if !boolObject(a.Left.Eval(scope)) {
		return builtin.BoolFalse
	}
	if boolObject(a.Right.Eval(scope)) {
		return builtin.BoolTrue
	}
	return builtin.BoolFalse
}

func (a *AndNode) Code() string {
	return fmt.Sprintf("%s and %s", a.Left.Code(), a.Right.Code())
}
