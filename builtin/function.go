package builtin

import (
	"fmt"
	"github.com/macbinn/hacklang/value"
)

type Function struct {
	Name *String
	fn func(...value.Object) value.Object
}

func (f *Function) Repr() string {
	return fmt.Sprintf("<function %s>", f.Name.S)
}

func (f *Function) Get(name string) value.Object {
	if name == "name" {
		return f.Name
	}
	return nil
}

func (f *Function) Func() func(...value.Object) value.Object {
	return f.fn
}

func NewFunction(name string, v func(...value.Object) value.Object) *Function {
	return &Function{
		Name: NewString(name),
		fn: v,
	}
}

func print(args ...value.Object) value.Object {
	for _, arg := range args {
		s, ok := arg.(*String)
		if ok {
			fmt.Printf("%s ", s.S)
		} else if arg == nil {
			fmt.Printf("<nil> ")
		} else {
			fmt.Printf("%s ", arg.Repr())
		}
	}
	fmt.Printf("\n")
	return nil
}

func sum(args ...value.Object) value.Object {
	sum := 0
	for _, a := range args {
		n := a.(*Number)
		sum += n.Int
	}
	return NewNumber(sum)
}

func getType(args ...value.Object) value.Object {
	obj := args[0]
	switch obj.(type) {
	case *String:
		return NewString("string")
	case *Number:
		return NewString("number")
	case *Bool:
		return NewString("bool")
	case *List:
		return NewString("list")
	case *Function:
		return NewString("function")
	case *Map:
		return NewString("map")
	}
	return NewString("unknown")
}

var (
	Print = NewFunction("print", print)
	Sum = NewFunction("sum", sum)
	Type = NewFunction("type", getType)
)

func init() {
	GlobalScope.Register("print", Print)
	GlobalScope.Register("sum", Sum)
	GlobalScope.Register("type", Type)
}