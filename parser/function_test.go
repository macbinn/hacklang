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
	}
	testHandler(t, functionHandler{}, cases)
}
