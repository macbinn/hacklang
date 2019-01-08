package parser

import (
	"testing"
)

func TestListHandler_Parse(t *testing.T) {
	cases := []testCase{
		{
			code: "[1, 2, 3] ",
			pos: 7,
			node: "<List Items=[<Number 1> <Number 2> <Number 3>]>",
		},
		{
			code: "[1, true, `hello`]",
			pos: 7,
			node: "<List Items=[<Number 1> <Bool true> <String hello>]>",
		},
		{
			code: "[1, true, `hello`, [], [1, 3]]",
			pos: 16,
			node: "<List Items=[<Number 1> <Bool true> <String hello> <List Items=[]> <List Items=[<Number 1> <Number 3>]>]>",
		},
		{
			code: "[]",
			pos: 2,
			node: "<List Items=[]>",
		},
		{
			code: "[1, []]",
			pos: 6,
			node: "<List Items=[<Number 1> <List Items=[]>]>",
		},
		{
			code: "[[[]]]",
			pos: 6,
			node: "<List Items=[<List Items=[<List Items=[]>]>]>",
		},
		{
			code: "[1, 2, ",
			err: ErrSyntaxError,
		},
		{
			code: "1, 2, 3",
			err: ErrSyntaxError,
		},
		{
			code: "",
			err: ErrSyntaxError,
		},
		{
			code: "[",
			err: ErrSyntaxError,
		},
	}
	testHandler(t, listHandler{}, cases)
}
