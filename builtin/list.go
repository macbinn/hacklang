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

var listMethods = map[string]interface{}{}

func (l *List) G(name string) *value.Object {
	//method, ok := listMethods[name]
	//if !ok {
	//	return nil
	//}
	//switch m := method.(type) {
	//case func(*list.List, func(...interface{})):
	//	return func(f func(...interface{})) {
	//		m(l.l, f)
	//	}
	//}
	return nil
}

func RegisterListMethod(name string, f interface{}) {
	listMethods[name] = f
}

func init() {
	RegisterListMethod("forEach", func(l *list.List, f func(...interface{})) {
		for i := l.Front(); i != nil; i = i.Next() {
			f(i.Value)
		}
	})
}
