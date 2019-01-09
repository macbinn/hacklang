package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/value"
	"log"
)

type IdNode struct {
	Name string
}

func (i *IdNode) String() string {
	return fmt.Sprintf("<Id Name=%s>", i.Name)
}

func (i *IdNode) Eval(scope *value.Scope) value.Object {
	v, ok := scope.Resolve(i.Name)
	if !ok {
		log.Printf("%s not found in scope", i.Name)
	}
	return v
}
