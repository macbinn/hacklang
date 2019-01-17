package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/value"
)

type DotNode struct {
	Left  Node
	Right Node
}

func (d *DotNode) Code() string {
	return fmt.Sprintf("%s.%s", d.Left.Code(), d.Right.Code())
}

func (d *DotNode) String() string {
	return fmt.Sprintf("<Dot Left=%s Right=%s>", d.Left, d.Right)
}

func (d *DotNode) Eval(scope *value.Scope) value.Object {
	left := d.Left.Eval(scope).(value.Getter)
	right := d.Right.(*IdNode)
	return left.Get(right.Name)
}
