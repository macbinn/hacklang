package parser

import (
	"fmt"
	"github.com/macbinn/hacklang/token"
	"testing"
)

type testCase struct {
	code string
	pos  int
	err  error
	node string
}

func testHandler(t *testing.T, handler Handler, cases []testCase) {
	for _, c := range cases {
		p := token.NewParser([]byte(c.code))
		tokens, err := p.Parse()
		if err != nil {
			t.Fatal(err)
		}
		node, pos, err := handler.Parse(tokens)
		if err != c.err {
			t.Errorf("%s expect err %v but got %v", c.code, c.err, err)
		} else if pos != c.pos {
			t.Errorf("%s expect pos %d but got %d", c.code, c.pos, pos)
		} else if c.err == nil && fmt.Sprintf("%s", node) != c.node {
			t.Errorf("%s except node %s but got %s", c.code, c.node, node)
		}
	}
}
