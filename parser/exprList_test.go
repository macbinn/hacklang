package parser

import "testing"

func TestExprListHandler_Parse(t *testing.T) {
	cases := []testCase{
		{
			code: "a = 1 print(a)",
			pos: 7,
			node: "<ExprList Nodes=[<Assign Left=<Id Name=a> Right=<Number 1>> <Call Callee=<Id Name=print> Arguments=[<Id Name=a>]>]>",
		},
		{
			code: "a = 1 b = 2 sum(a, b)",
			pos: 12,
			node: "<ExprList Nodes=[<Assign Left=<Id Name=a> Right=<Number 1>> <Assign Left=<Id Name=b> Right=<Number 2>> <Call Callee=<Id Name=sum> Arguments=[<Id Name=a> <Id Name=b>]>]>",
		},
		{
			code: "a = `abc` print(a.upper())",
			pos: 11,
			node: "<ExprList Nodes=[<Assign Left=<Id Name=a> Right=<String abc>> <Call Callee=<Id Name=print> Arguments=[<Call Callee=<Dot Left=<Id Name=a> Right=<Id Name=upper>> Arguments=[]>]>]>",
		},
		{
			code: "a = `abc` print(a.upper()) b = `ABC` print(b.lower())",
			pos: 22,
			node: "<ExprList Nodes=[<Assign Left=<Id Name=a> Right=<String abc>> <Call Callee=<Id Name=print> Arguments=[<Call Callee=<Dot Left=<Id Name=a> Right=<Id Name=upper>> Arguments=[]>]> <Assign Left=<Id Name=b> Right=<String ABC>> <Call Callee=<Id Name=print> Arguments=[<Call Callee=<Dot Left=<Id Name=b> Right=<Id Name=lower>> Arguments=[]>]>]>",
		},
	}

	testHandler(t, exprListHandler{}, cases)
}
