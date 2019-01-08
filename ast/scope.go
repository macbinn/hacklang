package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/builtin"
)

type Scope struct {
	Vars   *builtin.Map
	Parent *Scope
}

func (s *Scope) Get(name string) builtin.Object {
	return s.Vars.Get(name)
}

func (s *Scope) Repr() string {
	return fmt.Sprintf("<Scope %s>", s.Vars.Repr())
}

func NewScope(parent *Scope) *Scope {
	return &Scope{
		Vars:   builtin.NewEmptyMap(),
		Parent: parent,
	}
}

var GlobalScope = NewScope(nil)

// Resolve resolve name in scope
func (s *Scope) Resolve(name string) (builtin.Object, bool) {
	for scope := s; scope != nil; scope = s.Parent {
		v, ok := scope.Vars.Val[name]
		if ok {
			return v, true
		}
	}
	return nil, false
}

func (s *Scope) Register(name string, v builtin.Object) {
	s.Vars.Val[name] = v
}

func init() {
	GlobalScope.Register("global", GlobalScope)
	GlobalScope.Register("print", builtin.Print)
	GlobalScope.Register("sum", builtin.Sum)
	GlobalScope.Register("type", builtin.Type)
	GlobalScope.Register("web", builtin.WebExports)
	GlobalScope.Register("json", builtin.Json)
}
