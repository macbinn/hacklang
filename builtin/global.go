package builtin

import (
	"github.com/macbinn/hacklang/value"
)

var GlobalScope = value.NewScope(nil)

func init() {
	GlobalScope.Register("global", NewMap(GlobalScope.Vars))
}
