package ast

import "github.com/macbinn/hacklang/buildin"

type Node interface {
	Eval(scope *Scope) buildin.Object
}
