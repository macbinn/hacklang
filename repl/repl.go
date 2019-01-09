package repl

import (
	"bufio"
	"fmt"
	"github.com/macbinn/hacklang/builtin"
	"github.com/macbinn/hacklang/parser"
	"io"
	"log"
	"os"
)

func Repl(r io.Reader, w io.Writer) error {
	reader := bufio.NewReader(r)
	fmt.Fprintln(w, "hacklang v1.0 https://github.com/macbinn/hacklang")
	for {
		fmt.Fprintf(w, " -> ")
		input, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Fprintln(w, "\nBye!")
			return nil
		} else if err != nil {
			return err
		}
		// remove '\n' at the end
		input = input[:len(input)-1]
		if len(input) == 0 {
			continue
		}
		p := parser.NewParser([]byte(input))
		node, err := p.Parse()
		if err != nil {
			fmt.Fprintf(w, "error %s\n", err)
		} else {
			// log.Printf("%s\n", node)
			obj := node.Eval(builtin.GlobalScope)
			if obj != nil {
				fmt.Fprintf(w, "%s\n", obj.Repr())
			} else {
				fmt.Fprint(w, "<nil>\n")
			}
		}
	}
}

func Start() {
	err := Repl(os.Stdin, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
