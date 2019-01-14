package parser

import "testing"

func TestMapHandler_Parse(t *testing.T) {
	cases := []testCase {
		{
			code: "{a: 1, b: 2}",
			pos: 9,
			node: "<Map a=<Number 1> b=<Number 2>>",
		},
		{
			code: "{a: true}",
			pos: 5,
			node: "<Map a=<Bool true>>",
		},
		{
			code: "{}",
			pos: 2,
			node: "<Map>",
		},
	}
	testHandler(t, mapHandler{}, cases)
}
