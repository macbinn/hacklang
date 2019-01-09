package main

import (
	"github.com/macbinn/hacklang/repl"
	"github.com/macbinn/hacklang/require"
	"os"
)

func main() {
	if len(os.Args) == 1 { // no args
		repl.Start()
	} else {
		file := os.Args[1]
		require.ExecFile(file)
	}
}
