package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/buildin"
)

type Scope struct {
	Vars   *buildin.Map
	Parent *Scope
}

func (s *Scope) Get(name string) buildin.Object {
	return s.Vars.Get(name)
}

func (s *Scope) Repr() string {
	return fmt.Sprintf("<Scope %s>", s.Vars.Repr())
}

func NewScope(parent *Scope) *Scope {
	return &Scope{
		Vars:   buildin.NewEmptyMap(),
		Parent: parent,
	}
}

var GlobalScope = NewScope(nil)

// Resolve resolve name in scope
func (s *Scope) Resolve(name string) (buildin.Object, bool) {
	for scope := s; scope != nil; scope = s.Parent {
		v, ok := scope.Vars.Val[name]
		if ok {
			return v, true
		}
	}
	return nil, false
}

func (s *Scope) Register(name string, v buildin.Object) {
	s.Vars.Val[name] = v
}

func init() {
	GlobalScope.Register("global", GlobalScope)
	GlobalScope.Register("print", buildin.Print)
	GlobalScope.Register("sum", buildin.Sum)
	GlobalScope.Register("type", buildin.Type)
	GlobalScope.Register("web", buildin.WebExports)
	GlobalScope.Register("json", buildin.Json)
}
