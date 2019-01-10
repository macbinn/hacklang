package parser

import "testing"

func TestMapHandler_Parse(t *testing.T) {
	cases := []testCase {
		{
			code: "{}",
			pos: 2,
			node: "<Map>",
		},
	}
	testHandler(t, mapHandler{}, cases)
}
