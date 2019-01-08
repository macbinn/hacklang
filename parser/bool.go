package parser

import (
	"github.com/macbinn/hacklang/ast"
	"github.com/macbinn/hacklang/token"
)

type boolHandler struct {
}

func (boolHandler) Parse(tokens []*token.Token) (ast.Node, int, error) {
	if len(tokens) == 0 {
		return nil, 0, ErrSyntaxError
	}
	if tokens[0].Type != token.BOOL {
		return nil, 0, ErrSyntaxError
	}
	var value bool
	if tokens[0].Value == "true" {
		value = true
	} else if tokens[0].Value == "false" {
		value = false
	} else {
		return nil, 0, ErrSyntaxError
	}
	node := &ast.BoolNode{
		Value: value,
	}
	return node, 1, nil
}

func init() {
	Register("bool", boolHandler{})
}
