package parser

import "testing"

func TestExprHandler_Parse(t *testing.T) {
	cases := []testCase {
		{
			code: "123",
			pos: 1,
			node: "<Number 123>",
		},
		{
			code: "`abc`",
			pos: 1,
			node: "<String abc>",
		},
		{
			code: "true",
			pos: 1,
			node: "<Bool true>",
		},
		{
			code: "[1, 2, 3]",
			pos: 7,
			node: "<List Items=[<Number 1> <Number 2> <Number 3>]>",
		},
		{
			code: "f()",
			pos: 3,
			node: "<Call Callee=<Id Name=f> Arguments=[]>",
		},
		{
			code: "a",
			pos: 1,
			node: "<Id Name=a>",
		},
		{
			code: "i => {}",
			pos: 5,
			node: "<Function Arguments=[i], Body=<nil>>",
		},
		{
			code: "a = 1",
			pos: 3,
			node: "<Assign Left=<Id Name=a> Right=<Number 1>>",
		},
		{
			code: "a.b",
			pos: 3,
			node: "<Dot Left=<Id Name=a> Right=<Id Name=b>>",
		},
		{
			code: "print(a.b())",
			pos: 8,
			node: "<Call Callee=<Id Name=print> Arguments=[<Call Callee=<Dot Left=<Id Name=a> Right=<Id Name=b>> Arguments=[]>]>",
		},
	}
	testHandler(t, exprHandler{}, cases)
}
