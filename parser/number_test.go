package parser

import "testing"

func TestNumberHandler_Parse(t *testing.T) {
	cases := []testCase{
		{
			code: "123",
			pos:1,
			node: "<Number 123>",
		},
		{
			code: "",
			err: ErrSyntaxError,
		},
		{
			code: "abc",
			err: ErrSyntaxError,
		},
	}
	testHandler(t, numberHandler{}, cases)
}
