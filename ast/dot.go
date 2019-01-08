package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/builtin"
)

type DotNode struct {
	Left  Node
	Right Node
}

func (d *DotNode) String() string {
	return fmt.Sprintf("<Dot Left=%s Right=%s>", d.Left, d.Right)
}

func (d *DotNode) Eval(scope *Scope) builtin.Object {
	left := d.Left.Eval(scope)
	right := d.Right.(*IdNode)
	return left.Get(right.Name)
}
