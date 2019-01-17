package parser

import "testing"

func TestNotHandler_Parse(t *testing.T) {
	cases := []testCase {
		{
			code: "not true",
			pos: 2,
			node: "<Not Expr=<Bool true>>",
		},
		{
			code: "not false",
			pos: 2,
			node: "<Not Expr=<Bool false>>",
		},
		{
			code: "not 1",
			pos: 2,
			node: "<Not Expr=<Number 1>>",
		},
		{
			code: "not a()",
			pos: 4,
			node: "<Not Expr=<Call Callee=<Id Name=a> Arguments=[]>>",
		},
	}
	testHandler(t, notHandler{}, cases)
}
