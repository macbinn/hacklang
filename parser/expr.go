package parser

import (
	"github.com/macbinn/hacklang/ast"
	"github.com/macbinn/hacklang/token"
)

type exprHandler struct {
}

func (exprHandler) Parse(tokens []*token.Token) (ast.Node, int, error) {
	expr, pos, err := ParseGreedy(tokens,
		"number",
		"string",
		"bool",
		"list",
		"map",
		"call",
		"id",
		"function",
		"assign",
		"dot",
		"if",
	)
	if err != nil {
		return nil, 0, ErrSyntaxError
	}
	return expr, pos, nil
}

func init() {
	Register("expr", exprHandler{})
}
