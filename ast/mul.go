package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/builtin"
	"github.com/macbinn/hacklang/value"
)

type Mul struct {
	Left  Node
	Right Node
}

func (m *Mul) String() string {
	return fmt.Sprintf("<Mul Left=%s Right=%s>", m.Left, m.Right)
}

func (a *Mul) Eval(scope *value.Scope) value.Object {
	left := a.Left.Eval(scope).(*builtin.Number).Int
	right := a.Right.Eval(scope).(*builtin.Number).Int
	return builtin.NewNumber(left * right)
}

func (a *Mul) Code() string {
	return fmt.Sprintf("%s + %s", a.Left.Code(), a.Right.Code())
}
