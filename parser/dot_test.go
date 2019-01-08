package parser

import "testing"

func TestDotHandler_Parse(t *testing.T) {
	cases := []testCase{
		{
			code: "a.a",
			pos: 3,
			err: nil,
			node: "<Dot Left=<Id Name=a> Right=<Id Name=a>>",
		},
		{
			code: "`abc`.a",
			pos: 3,
			err: nil,
			node: "<Dot Left=<String abc> Right=<Id Name=a>>",
		},
		{
			code: "[1, 2, 3].a",
			pos: 9,
			err: nil,
			node: "<Dot Left=<List Items=[<Number 1> <Number 2> <Number 3>]> Right=<Id Name=a>>",
		},
		{
			code: "f().a",
			pos: 5,
			err: nil,
			node: "<Dot Left=<Call Callee=<Id Name=f> Arguments=[]> Right=<Id Name=a>>",
		},
		{
			code: "a.b.c.e.f",
			pos: 9,
			err: nil,
			node: "<Dot Left=<Dot Left=<Dot Left=<Dot Left=<Id Name=a> Right=<Id Name=b>> Right=<Id Name=c>> Right=<Id Name=e>> Right=<Id Name=f>>",
		},
		{
			code: "a b c d e",
			pos: 0,
			err: ErrSyntaxError,
		},
	}
	testHandler(t, dotHandler{}, cases)
}
