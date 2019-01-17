package parser

import "testing"

func TestAndHandler_Parse(t *testing.T) {
	cases := []testCase {
		{
			code: "true and true",
			pos: 3,
			node: "<And Left=<Bool true> Right=<Bool true>>",
		},
		{
			code: "1 and false",
			pos: 3,
			node: "<And Left=<Number 1> Right=<Bool false>>",
		},
		{
			code: "a() and b()",
			pos: 7,
			node: "<And Left=<Call Callee=<Id Name=a> Arguments=[]> Right=<Call Callee=<Id Name=b> Arguments=[]>>",
		},
	}
	testHandler(t, exprHandler{}, cases)
}
