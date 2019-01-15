package parser

import "testing"

func TestFunctionHandler_Parse(t *testing.T) {
	cases := []testCase{
		{
			code: "i => {}",
			pos: 5,
			node: "<Function Arguments=[i], Body=[]>",
		},
		{
			code: "i => { f() }",
			pos: 8,
			node: "<Function Arguments=[i], Body=[<Call Callee=<Id Name=f> Arguments=[]>]>",
		},
		{
			code: "i => { 1 }",
			pos: 6,
			node: "<Function Arguments=[i], Body=[<Number 1>]>",
		},
		{
			code: "(a, b, c) => {}",
			pos: 11,
			node: "<Function Arguments=[a b c], Body=[]>",
		},
		{
			code: "i => { a = {} print(a) }",
			pos: 13,
			node: "<Function Arguments=[i], Body=[<Assign Left=<Id Name=a> Right=<Map>> <Call Callee=<Id Name=print> Arguments=[<Id Name=a>]>]>",
		},
		{
			code: "l => { l.forEach(i => { print(i) }) }",
			pos: 19,
			node: "<Function Arguments=[l], Body=[<Call Callee=<Dot Left=<Id Name=l> Right=<Id Name=forEach>> Arguments=[<Function Arguments=[i], Body=[<Call Callee=<Id Name=print> Arguments=[<Id Name=i>]>]>]>]>",
		},
	}
	testHandler(t, functionHandler{}, cases)
}
