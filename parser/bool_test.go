package parser

import "testing"

func TestBoolHandler_Parse(t *testing.T) {
	cases := []testCase {
		{
			code: "true",
			pos: 1,
			node: "<Bool true>",
		},
		{
			code: "false",
			pos: 1,
			node: "<Bool false>",
		},
		{
			code: "123",
			err: ErrSyntaxError,
		},
		{
			code: "",
			err: ErrSyntaxError,
		},
		{
			code: "trueA",
			err: ErrSyntaxError,
		},
		{
			code: "falseB",
			err: ErrSyntaxError,
		},
	}

	testHandler(t, boolHandler{}, cases)
}
