package parser

import "testing"

func TestCallHandler_Parse(t *testing.T) {
	cases := []testCase{
		{
			code: "f()",
			pos: 3,
			node: "<Call Callee=<Id Name=f> Arguments=[]>",
		},
		{
			code: "f(1)",
			pos: 4,
			node: "<Call Callee=<Id Name=f> Arguments=[<Number 1>]>",
		},
		{
			code: "f(1, 2, 3)",
			pos: 8,
			node: "<Call Callee=<Id Name=f> Arguments=[<Number 1> <Number 2> <Number 3>]>",
		},
		{
			code: "a.b(1)",
			pos: 6,
			node: "<Call Callee=<Dot Left=<Id Name=a> Right=<Id Name=b>> Arguments=[<Number 1>]>",
		},
		{
			code: "i => { print(i) }()",
			pos: 11,
			node: "<Call Callee=<Function Arguments=[i], Body=<ExprList Nodes=[<Call Callee=<Id Name=print> Arguments=[<Id Name=i>]>]>> Arguments=[]>",
		},
		{
			code:"print(a.b())",
			pos: 8,
			err: nil,
			node: "<Call Callee=<Id Name=print> Arguments=[<Call Callee=<Dot Left=<Id Name=a> Right=<Id Name=b>> Arguments=[]>]>",
		},
		{
			code:"print(a.b()) b = `ABC` print(b.lower())",
			pos: 8,
			err: nil,
			node: "<Call Callee=<Id Name=print> Arguments=[<Call Callee=<Dot Left=<Id Name=a> Right=<Id Name=b>> Arguments=[]>]>",
		},
		{
			code: "f()()",
			pos: 5,
			node: "<Call Callee=<Call Callee=<Id Name=f> Arguments=[]> Arguments=[]>",
		},
		// todo: make it work
		//{
		//	code: "`hello`.upper().lower()",
		//	pos: 9,
		//},
	}
	testHandler(t, callHandler{}, cases)
}
