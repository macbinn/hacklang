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
		{
			code: `{
  a: 1,
  b: l => {
    l.forEach(i => {
      print(i)
    })
  }
}`,
			pos: 27,
			node: "<Map a=<Number 1> b=<Function Arguments=[l], Body=[<Call Callee=<Dot Left=<Id Name=l> Right=<Id Name=forEach>> Arguments=[<Function Arguments=[i], Body=[<Call Callee=<Id Name=print> Arguments=[<Id Name=i>]>]>]>]>>",
		},
	}
	testHandler(t, mapHandler{}, cases)
}
