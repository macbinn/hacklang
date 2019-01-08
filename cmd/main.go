package main

import (
	"github.com/macbinn/hacklang/ast"
	"github.com/macbinn/hacklang/parser"
	"github.com/macbinn/hacklang/repl"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 { // no args
		repl.Start()
	} else {
		file := os.Args[1]
		code, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		p := parser.NewParser(code)
		node, err := p.Parse()
		if err != nil {
			log.Fatal(err)
		}
		//log.Printf("%s", node)
		node.Eval(ast.GlobalScope)
	}
}
