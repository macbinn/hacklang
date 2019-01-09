package builtin

import (
	"fmt"
	"github.com/macbinn/hacklang/value"
)

type Number struct {
	Int int
}

func (n *Number) Repr() string {
	return fmt.Sprintf("%d", n.Int)
}

func (n *Number) Get(name string) value.Object {
	return nil
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

func (*Bool) Get(name string) value.Object {
	return nil
}

func NewBool(v bool) *Bool {
	return &Bool{
		Val: v,
	}
}
