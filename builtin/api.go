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
	if v {
		return BoolTrue
	}
	return BoolFalse
}

var (
	BoolTrue  = &Bool{true}
	BoolFalse = &Bool{false}
)
