package ast

import "github.com/macbinn/hacklang/builtin"

type Node interface {
	Eval(scope *Scope) builtin.Object
}
