package parser

import (
	"github.com/macbinn/hacklang/ast"
	"github.com/macbinn/hacklang/token"
	"strconv"
)

type numberHandler struct {
}

func (numberHandler) Parse(tokens []*token.Token) (ast.Node, int, error) {
	if len(tokens) < 1 || tokens[0].Type != token.NUMBER {
		return nil, 0, ErrSyntaxError
	}
	value, err := strconv.Atoi(tokens[0].Value)
	if err != nil {
		return nil, 0, ErrSyntaxError
	}
	node := &ast.NumberNode{
		Value: value,
	}
	return node, 1, nil
}

func init() {
	Register("number", numberHandler{})
}
