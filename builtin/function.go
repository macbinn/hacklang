package builtin

import (
	"fmt"
)

type Function struct {
	Name *String
	fn func(...Object) Object
}

func (f *Function) Repr() string {
	return fmt.Sprintf("<function %s>", f.Name.s)
}

func (f *Function) Get(name string) Object {
	if name == "name" {
		return f.Name
	}
	return nil
}

func (f *Function) Func() func(...Object) Object {
	return f.fn
}

func NewFunction(name string, v func(...Object) Object) *Function {
	return &Function{
		Name: NewString(name),
		fn: v,
	}
}

func print(args ...Object) Object {
	for _, arg := range args {
		s, ok := arg.(*String)
		if ok {
			fmt.Printf("%s ", s.s)
		} else if args == nil {
			fmt.Printf("<nil> ")
		} else {
			fmt.Printf("%s ", arg.Repr())
		}
	}
	fmt.Printf("\n")
	return nil
}

func sum(args ...Object) Object {
	sum := 0
	for _, a := range args {
		n := a.(*Number)
		sum += n.Int
	}
	return NewNumber(sum)
}

func getType(args ...Object) Object {
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
