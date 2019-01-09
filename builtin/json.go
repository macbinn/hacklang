package builtin

import (
	"encoding/json"
	"github.com/macbinn/hacklang/value"
)

func Convert(object value.Object) interface{} {
	switch obj := object.(type) {
	case *String:
		return obj.S
	case *Number:
		return obj.Int
	case *Bool:
		return obj.Val
	case *Map:
		m := map[string]interface{}{}
		for name, value := range obj.Val {
			m[name] = Convert(value)
		}
		return m
	case *List:
		l := make([]interface{}, 0)
		for i := obj.L.Front(); i != nil; i = i.Next() {
			l = append(l, Convert(i.Value.(value.Object)))
		}
		return l
	}
	return nil
}

func Obj(v interface{}) value.Object {
	switch o := v.(type) {
	case string:
		return NewString(o)
	case int:
		return NewNumber(o)
	case bool:
		return NewBool(o)
	case map[string]interface{}:
		m := NewEmptyMap()
		for name, value := range o {
			m.Val[name] = Obj(value)
		}
		return m
	}
	return nil
}

func JsonEncode(args ...value.Object) value.Object {
	bs, err := json.Marshal(Convert(args[0]))
	if err != nil {
		return nil
	}
	return NewString(string(bs))
}

func JsonDecode(args ...value.Object) value.Object {
	s := args[0].(*String).S
	var v map[string]interface{}
	err := json.Unmarshal([]byte(s), &v)
	if err != nil {
		return nil
	}
	return Obj(v)
}

var (
	Json = NewMap(map[string]value.Object{
		"encode": NewFunction("json.encode", JsonEncode),
		"decode": NewFunction("json.decode", JsonDecode),
	})
)

func init() {
	GlobalScope.Register("json", Json)
}