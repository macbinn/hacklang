package ast

import (
	"fmt"
	"github.com/macbinn/hacklang/builtin"
	"github.com/macbinn/hacklang/value"
)

type IfNode struct {
	Condition Node
	Body Node
}

func (i *IfNode) String() string {
	return fmt.Sprintf("<If Condition=%v Body=%v>", i.Condition, i.Body)
}

func boolObject(object value.Object) bool {
	switch o := object.(type) {
	case *builtin.Bool:
		return o.Val
	case *builtin.Number:
		return o.Int > 0
	case *builtin.String:
		return len(o.S) > 0
	case *builtin.List:
		return o.L.Len() > 0
	case *builtin.Map:
		return len(o.Val) > 0
	}
	return false
}

func (i *IfNode) Eval(scope *value.Scope) value.Object {
	cond := i.Condition.Eval(scope)
	var obj value.Object
	if boolObject(cond) {
		obj = i.Body.Eval(scope)
	}
	return obj
}

func (i *IfNode) Code() string {
	return fmt.Sprintf("if %s {\n%s\n}", i.Condition.Code(), i.Body.Code())
}



