package parser

import "testing"

func TestReturnHandler_Parse(t *testing.T) {
	cases := []testCase {
		{
			code: "return 1",
			pos: 2,
			node: "<Return Expr=<Number 1>>",
		},
		{
			code: "return true",
			pos: 2,
			node: "<Return Expr=<Bool true>>",
		},
		{
			code: "return",
			err: ErrSyntaxError,
		},
	}
	testHandler(t, returnHandler{}, cases)
}
