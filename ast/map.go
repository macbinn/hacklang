package ast

import (
	"github.com/macbinn/hacklang/builtin"
	"github.com/macbinn/hacklang/value"
)

type MapNode struct {

}

func (m *MapNode) String() string {
	return "<Map>"
}

func (m *MapNode) Code() string {
	return "{}"
}

func (m *MapNode) Eval(scope *value.Scope) value.Object {
	return builtin.NewEmptyMap()
}
