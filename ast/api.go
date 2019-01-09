package ast

import "github.com/macbinn/hacklang/value"

type Node interface {
	Eval(scope *value.Scope) value.Object
	//Code() string
}
