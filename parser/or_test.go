package parser

import "testing"

func TestOrHandler_Parse(t *testing.T) {
	cases := []testCase {
		{
			code: "true or true",
			pos: 3,
			node: "<Or Left=<Bool true> Right=<Bool true>>",
		},
		{
			code: "1 or false",
			pos: 3,
			node: "<Or Left=<Number 1> Right=<Bool false>>",
		},
		{
			code: "a() or b()",
			pos: 7,
			node: "<Or Left=<Call Callee=<Id Name=a> Arguments=[]> Right=<Call Callee=<Id Name=b> Arguments=[]>>",
		},
	}
	testHandler(t, exprHandler{}, cases)
}
