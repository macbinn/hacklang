package builtin

import (
	"fmt"
	"strconv"
	"strings"
)

type String struct {
	s string
}

func (s *String) Repr() string {
	return fmt.Sprintf("`%s`", s.s)
}

func (s *String) Get(name string) Object {
	method := s.GetMethod(name)
	if method != nil {
		return NewFunction("string." + name, method)
	}
	return nil
}

func NewString(s string) *String {
	return &String{
		s: s,
	}
}

var stringMethods = map[string]interface{}{}

func RegisterStringMethod(name string, f interface{}) {
	stringMethods[name] = f
}

func (s *String) GetMethod(name string) func(...Object) Object {
	method, ok := stringMethods[name]
	if !ok {
		return nil
	}
	switch m := method.(type) {
	case func(string) string:
		return func(...Object) Object {
			return NewString(m(s.s))
		}
	case func(string) int:
		return func(...Object) Object {
			return NewNumber(m(s.s))
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
