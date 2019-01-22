package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/builtin"
	"github.com/macbinn/hacklang/value"
)

type Dev struct {
	Left  Node
	Right Node
}

func (d *Dev) String() string {
	return fmt.Sprintf("<Dev Left=%s Right=%s>", d.Left, d.Right)
}

func (d *Dev) Eval(scope *value.Scope) value.Object {
	left := d.Left.Eval(scope).(*builtin.Number).Int
	right := d.Right.Eval(scope).(*builtin.Number).Int
	return builtin.NewNumber(left / right)
}

func (d *Dev) Code() string {
	return fmt.Sprintf("%s / %s", d.Left.Code(), d.Right.Code())
}
