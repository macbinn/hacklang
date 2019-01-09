package require

import (
	"github.com/macbinn/hacklang/builtin"
	"github.com/macbinn/hacklang/parser"
	"github.com/macbinn/hacklang/value"
	"io/ioutil"
	"log"
	"strings"
)

const FileExt = ".hack"

func resolveFile(file string) string {
	if !strings.HasSuffix(file, FileExt) {
		return file + FileExt
	}
	return file
}

func ExecFile(file string) value.Object {
	code, err := ioutil.ReadFile(resolveFile(file))
	if err != nil {
		log.Fatal(err)
	}
	p := parser.NewParser(code)
	node, err := p.Parse()
	if err != nil {
		log.Fatal(err)
	}
	return node.Eval(builtin.GlobalScope)
}

// require file
// sql = require(`sql`)
// sql.connect(`xxx`)
func require(args ...value.Object) value.Object {
	fn := args[0].(*builtin.String).S
	return ExecFile(fn)
}

var (
	Require = builtin.NewFunction("require", require)
)

func init() {
	builtin.GlobalScope.Register("require", Require)
}