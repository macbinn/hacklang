package builtin

import (
	"fmt"
)

type Number struct {
	Int int
}

func (n *Number) Repr() string {
	return fmt.Sprintf("%d", n.Int)
}

func NewNumber(v int) *Number {
	return &Number{
		Int: v,
	}
}

type Bool struct {
	Val bool
}

func (b *Bool) Repr() string {
	if b.Val {
		return "true"
	}
	return "false"
}

func NewBool(v bool) *Bool {
	return &Bool{
		Val: v,
	}
}

var (
	BoolTrue = NewBool(true)
	BoolFalse = NewBool(false)
)
