package builtin

import (
	"container/list"
	"github.com/macbinn/hacklang/value"
	"strings"
)

type List struct {
	L *list.List
}

func (l *List) Get(name string) value.Object {
	switch name {
	case "forEach":
		return NewFunction("list.forEach", func(args ...value.Object) value.Object {
			return l.forEach(args...)
		})
	}
	return nil
}

func (l *List) Repr() string {
	builder := new(strings.Builder)
	builder.WriteByte('[')
	for i := l.L.Front(); i != nil; i = i.Next() {
		obj := i.Value.(value.Object)
		builder.WriteString(obj.Repr())
		if i.Next() != nil {
			builder.WriteString(", ")
		}
	}
	builder.WriteByte(']')
	return builder.String()
}

func NewEmptyList() *List {
	return &List{
		L: list.New(),
	}
}

func NewList(l *list.List) *List {
	return &List{L: l}
}

func (l *List) forEach(args...value.Object) value.Object {
	f := args[0].(*Function).fn
	for i := l.L.Front(); i != nil; i = i.Next() {
		obj := i.Value.(value.Object)
		f(obj)
	}
	return nil
}
