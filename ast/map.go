package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/builtin"
	"github.com/macbinn/hacklang/value"
	"strings"
)

type MapNode struct {
	Init map[string]Node
}

func (m *MapNode) String() string {
	if len(m.Init) == 0 {
		return "<Map>"
	}
	var items []string
	for name, node := range m.Init {
		items = append(items, fmt.Sprintf("%s=%s", name, node))
	}
	return fmt.Sprintf("<Map %s>", strings.Join(items, " "))
}

func (m *MapNode) Code() string {
	var items []string
	for name, node := range m.Init {
		items = append(items, fmt.Sprintf("%s: %s", name, node.Code()))
	}
	return fmt.Sprintf("{\n%s\n}", strings.Join(items, "\n"))
}

func (m *MapNode) Eval(scope *value.Scope) value.Object {
	ma := builtin.NewEmptyMap()
	for name, node := range m.Init {
		ma.Val[name] = node.Eval(scope)
	}
	return ma
}
