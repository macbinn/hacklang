package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/builtin"
	"github.com/macbinn/hacklang/value"
)

type Add struct {
	Left  Node
	Right Node
}

func (a *Add) String() string {
	return fmt.Sprintf("<Add Left=%s Right=%s>", a.Left, a.Right)
}

func (a *Add) Eval(scope *value.Scope) value.Object {
	left := a.Left.Eval(scope).(*builtin.Number).Int
	right := a.Right.Eval(scope).(*builtin.Number).Int
	return builtin.NewNumber(left + right)
}

func (a *Add) Code() string {
	return fmt.Sprintf("%s + %s", a.Left.Code(), a.Right.Code())
}
