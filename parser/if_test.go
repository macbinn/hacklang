package parser

import "testing"

func TestIfHandler_Parse(t *testing.T) {
	cases := []testCase {
		{
			code: "if true { print(1) }",
			pos: 8,
			node: "<If Condition=<Bool true> Body=<ExprList Nodes=[<Call Callee=<Id Name=print> Arguments=[<Number 1>]>]>>",
		},
		{
			code: "if true { a = 1 f(a) }",
			pos: 11,
			node: "<If Condition=<Bool true> Body=<ExprList Nodes=[<Assign Left=<Id Name=a> Right=<Number 1>> <Call Callee=<Id Name=f> Arguments=[<Id Name=a>]>]>>",
		},
	}
	testHandler(t, ifHandler{}, cases)
}