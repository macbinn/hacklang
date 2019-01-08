package parser

import "testing"

func TestStringHandler_Parse(t *testing.T) {
	cases := []testCase{
		{
			code: "`1234`",
			pos: 1,
			node: "<String 1234>",
		},
		{
			code: "",
			err: ErrSyntaxError,
		},
	}
	testHandler(t, stringHandler{}, cases)
}
