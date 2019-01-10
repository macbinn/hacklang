package main

import (
	"flag"
	"fmt"
	"github.com/macbinn/hacklang/format"
	"github.com/macbinn/hacklang/repl"
	"github.com/macbinn/hacklang/require"
	"log"
)

var (
	Format = flag.Bool("format", false, "format code")
)

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 { // no args
		repl.Start()
	} else {
		file := args[0]
		if *Format {
			code, err := format.Format(file)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Print(code)
		} else {
			require.ExecFile(file)
		}
	}
}
