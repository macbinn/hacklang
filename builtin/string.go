package builtin

import (
	"fmt"
	"github.com/macbinn/hacklang/value"
	"strconv"
	"strings"
)

type String struct {
	S string
}

func (s *String) Repr() string {
	return fmt.Sprintf("`%s`", s.S)
}

func (s *String) Get(name string) value.Object {
	method := s.GetMethod(name)
	if method != nil {
		return NewFunction("string." + name, method)
	}
	return nil
}

func NewString(s string) *String {
	return &String{
		S: s,
	}
}

var stringMethods = map[string]interface{}{}

func RegisterStringMethod(name string, f interface{}) {
	stringMethods[name] = f
}

func (s *String) GetMethod(name string) func(...value.Object) value.Object {
	method, ok := stringMethods[name]
	if !ok {
		return nil
	}
	switch m := method.(type) {
	case func(string) string:
		return func(...value.Object) value.Object {
			return NewString(m(s.S))
		}
	case func(string) int:
		return func(...value.Object) value.Object {
			return NewNumber(m(s.S))
		}
	}
	return nil
}

func init() {
	RegisterStringMethod("upper", strings.ToUpper)
	RegisterStringMethod("lower", strings.ToLower)
	RegisterStringMethod("trim", strings.TrimSpace)
	RegisterStringMethod("toNumber", func(s string) int {
		i, err := strconv.Atoi(s)
		if err != nil {
			return 0
		}
		return i
	})
}
