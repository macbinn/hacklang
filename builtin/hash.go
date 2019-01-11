package builtin

import (
	"github.com/macbinn/hacklang/value"
	"golang.org/x/crypto/bcrypt"
)

func BcryptGenerate(args...value.Object) value.Object {
	s := args[0].(*String).S
	bs, err := bcrypt.GenerateFromPassword([]byte(s), 10)
	if err != nil {
		return nil
	}
	return NewString(string(bs))
}

func BcryptCompare(args...value.Object) value.Object {
	hash := args[0].(*String).S
	pass := args[1].(*String).S
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err == nil {
		return NewBool(true)
	}
	return NewBool(false)
}

func init() {
	GlobalScope.Register("hash", NewMap(map[string]value.Object{
		"bcrypt": NewMap(map[string]value.Object{
			"generate": NewFunction("bcrypt.generate", BcryptGenerate),
			"compare": NewFunction("bcrypt.compare", BcryptCompare),
		}),
	}))
}