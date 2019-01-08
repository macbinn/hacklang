package parser

import "testing"

func TestAssignHandler_Parse(t *testing.T) {
	cases := []testCase {
		{
			code: "a = 1",
			pos: 3,
			node: "<Assign Left=<Id Name=a> Right=<Number 1>>",
		},
		{
			code: "a = `abc`",
			pos: 3,
			node: "<Assign Left=<Id Name=a> Right=<String abc>>",
		},
		{
			code: "a = true",
			pos: 3,
			node: "<Assign Left=<Id Name=a> Right=<Bool true>>",
		},
		{
			code: "a = [1, 2, 3]",
			pos: 9,
			node: "<Assign Left=<Id Name=a> Right=<List Items=[<Number 1> <Number 2> <Number 3>]>>",
		},
		{
			code: "a = f()",
			pos: 5,
			node: "<Assign Left=<Id Name=a> Right=<Call Callee=<Id Name=f> Arguments=[]>>",
		},
		{
			code: "a = b = 1",
			pos: 5,
			node: "<Assign Left=<Id Name=a> Right=<Assign Left=<Id Name=b> Right=<Number 1>>>",
		},
		{
			code: "",
			err: ErrSyntaxError,
		},
		{
			code: "i => {}",
			err: ErrSyntaxError,
		},
	}

	testHandler(t, assignHandler{}, cases)
}
