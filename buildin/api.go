package buildin

import "fmt"

// Object is the basic value in the hacklang world
type Object interface {
	// everything has `getter`
	Get(name string) Object
	// how to print in repl
	Repr() string
}

type Number struct {
	Int int
}

func (n *Number) Repr() string {
	return fmt.Sprintf("%d", n.Int)
}

func (n *Number) Get(name string) Object {
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

func (*Bool) Get(name string) Object {
	return nil
}

func NewBool(v bool) *Bool {
	return &Bool{
		Val: v,
	}
}
