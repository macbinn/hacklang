package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/value"
)

type AssignNode struct {
	Left  Node
	Right Node
}

func (a *AssignNode) Code() string {
	return fmt.Sprintf("%s = %s", a.Left.Code(), a.Right.Code())
}

func (a *AssignNode) String() string {
	return fmt.Sprintf("<Assign Left=%v Right=%v>", a.Left, a.Right)
}

func (a *AssignNode) Eval(scope *value.Scope) value.Object {
	if a.Right == nil {
		return nil
	}
	val := a.Right.Eval(scope)
	switch left := a.Left.(type) {
	case *IdNode:
		scope.Register(left.Name, val)
	case *DotNode:
		obj := left.Left.Eval(scope).(value.Setter)
		name := left.Right.(*IdNode).Name
		obj.Set(name, val)
	}
	return val
}
