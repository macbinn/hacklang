package builtin

import (
	"container/list"
	"strings"
)

type List struct {
	l *list.List
}

func (l *List) Get(name string) Object {
	return nil
}

func (l *List) Repr() string {
	builder := new(strings.Builder)
	builder.WriteByte('[')
	for i := l.l.Front(); i != nil; i = i.Next() {
		obj := i.Value.(Object)
		builder.WriteString(obj.Repr())
		if i.Next() != nil {
			builder.WriteString(", ")
		}
	}
	builder.WriteByte(']')
	return builder.String()
}

func NewList(l *list.List) *List {
	return &List{l: l}
}

var listMethods = map[string]interface{}{}

func (l *List) G(name string) *Object {
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
