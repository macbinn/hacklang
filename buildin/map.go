package buildin

import "strings"

type Map struct {
	Val map[string]Object
}

func (m *Map) Get(name string) Object {
	return m.Val[name]
}

func (m *Map) Repr() string {
	builder := new(strings.Builder)
	builder.WriteString("{")
	i := len(m.Val)
	for name, val := range m.Val {
		if name == "global" {
			i --
			continue
		}
		builder.WriteString(name)
		builder.WriteString(": ")
		obj := val.(Object)
		builder.WriteString(obj.Repr())
		i --
		if i > 0 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("}")
	return builder.String()
}

func NewMap(v map[string]Object) *Map {
	return &Map{
		Val: v,
	}
}

func NewEmptyMap() *Map {
	return &Map{
		Val: map[string]Object{},
	}
}

