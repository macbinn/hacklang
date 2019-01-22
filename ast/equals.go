package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/builtin"
	"github.com/macbinn/hacklang/value"
)

type Equals struct {
	Left  Node
	Right Node
}

func (a *Equals) String() string {
	return fmt.Sprintf("<Equals Left=%s Right=%s>", a.Left, a.Right)
}

func (a *Equals) Eval(scope *value.Scope) value.Object {
	switch left := a.Left.Eval(scope).(type) {
	case *builtin.Number:
		right := a.Right.Eval(scope).(*builtin.Number)
		return builtin.NewBool(left.Int == right.Int)
	case *builtin.Bool:
		right := a.Right.Eval(scope).(*builtin.Bool)
		return builtin.NewBool(left.Val == right.Val)
	case *builtin.String:
		right := a.Right.Eval(scope).(*builtin.String)
		return builtin.NewBool(left.S == right.S)
	}
	return nil
}

func (a *Equals) Code() string {
	return fmt.Sprintf("%s + %s", a.Left.Code(), a.Right.Code())
}
