package repl

import (
	"bytes"
	"strings"
	"testing"
)

func TestRepl(t *testing.T) {

	banner := "hacklang v1.0 https://github.com/macbinn/hacklang"

	cases := []struct {
		in  string
		out string
	}{
		{
			in: "1",
			out: "1",
		},
		{
			in: "`hello`",
			out: "`hello`",
		},
		{
			in: "true",
			out: "true",
		},
		{
			in: "a",
			out: "<nil>",
		},
		{
			in: "a = 1 a",
			out: "1",
		},
		{
			in: "`hello`.upper()",
			out: "`HELLO`",
		},
		{
			in: "type(1)",
			out: "`number`",
		},
		{
			in: "type(`hello`)",
			out: "`string`",
		},
		{
			in: "type([1, 2])",
			out: "`list`",
		},
		{
			in: "type(true)",
			out: "`bool`",
		},
		{
			in: "type(type)",
			out: "`function`",
		},
		{
			in: "type(web)",
			out: "`map`",
		},
		{
			in: "sum(1, 2)",
			out: "3",
		},
		{
			in: "sum(1, 2, 3, 4)",
			out: "10",
		},
		{
			in: "[1, 2, 3]",
			out: "[1, 2, 3]",
		},
		{
			in: "i => { print(1) }",
			out: "<function f>",
		},
		{
			in: "json.encode(1)",
			out: "`1`",
		},
		{
			in: "json.encode(`hello`)",
			out: "`\"hello\"`",
		},
		{
			in: "json.encode(true)",
			out: "`true`",
		},
		{
			in: "json.encode([1, 2, 3])",
			out: "`[1,2,3]`",
		},
		{
			in: "{}",
			out: "{}",
		},
		{
			in: "a = {} a.a = 1 a.a",
			out: "1",
		},
		// todo: make it work
		//{"`hello`.upper().lower()\n", "hello\n -> "},
	}

	for _, c := range cases {
		r := strings.NewReader(c.in + "\n")
		buf := new(bytes.Buffer)
		err := Repl(r, buf)
		if err != nil {
			t.Error(err)
		}
		expect := banner + "\n -> " + c.out + "\n -> \nBye!\n"
		got := buf.String()
		if expect != got {
			t.Errorf("input %s except %s but got %s", c.in, expect, got)
		}
	}
}
