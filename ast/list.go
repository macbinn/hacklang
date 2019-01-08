package ast

import (
	"container/list"
	"fmt"
	"github.com/macbinn/hacklang/buildin"
)

type ListNode struct {
	Items []Node
}

func (l *ListNode) String() string {
	return fmt.Sprintf("<List Items=%v>", l.Items)
}

func (l *ListNode) Eval(scope *Scope) buildin.Object {
	li := list.New()
	for _, node := range l.Items {
		li.PushBack(node.Eval(scope))
	}
	return buildin.NewList(li)
}
