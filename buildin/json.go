package buildin

import (
	"encoding/json"
)

func Convert(object Object) interface{} {
	switch obj := object.(type) {
	case *String:
		return obj.s
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
		for i := obj.l.Front(); i != nil; i = i.Next() {
			l = append(l, Convert(i.Value.(Object)))
		}
		return l
	}
	return nil
}

func Obj(v interface{}) Object {
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

func JsonEncode(args ...Object) Object {
	bs, err := json.Marshal(Convert(args[0]))
	if err != nil {
		return nil
	}
	return NewString(string(bs))
}

func JsonDecode(args ...Object) Object {
	s := args[0].(*String).s
	var v map[string]interface{}
	err := json.Unmarshal([]byte(s), &v)
	if err != nil {
		return nil
	}
	return Obj(v)
}

var (
	Json = NewMap(map[string]Object{
		"encode": NewFunction("json.encode", JsonEncode),
		"decode": NewFunction("json.decode", JsonDecode),
	})
)