package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/builtin"
)

type AssignNode struct {
	Left  Node
	Right Node
}

func (a *AssignNode) String() string {
	return fmt.Sprintf("<Assign Left=%v Right=%v>", a.Left, a.Right)
}

func (a *AssignNode) Eval(scope *Scope) builtin.Object {
	if a.Right == nil {
		return nil
	}
	value := a.Right.Eval(scope)
	switch left := a.Left.(type) {
	case *IdNode:
		scope.Register(left.Name, value)
	}
	return value
}
