package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/builtin"
	"github.com/macbinn/hacklang/value"
)

type Min struct {
	Left  Node
	Right Node
}

func (m *Min) String() string {
	return fmt.Sprintf("<Min Left=%s Right=%s>", m.Left, m.Right)
}

func (m *Min) Eval(scope *value.Scope) value.Object {
	left := m.Left.Eval(scope).(*builtin.Number).Int
	right := m.Right.Eval(scope).(*builtin.Number).Int
	return builtin.NewNumber(left - right)
}

func (m *Min) Code() string {
	return fmt.Sprintf("%s - %s", m.Left.Code(), m.Right.Code())
}
