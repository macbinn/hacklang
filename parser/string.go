package parser

import (
	"github.com/macbinn/hacklang/ast"
	"github.com/macbinn/hacklang/token"
)

type stringHandler struct {
}

func (stringHandler) Parse(tokens []*token.Token) (ast.Node, int, error) {
	if len(tokens) < 1 || tokens[0].Type != token.STRING {
		return nil, 0, ErrSyntaxError
	}
	node := &ast.StringNode{
		Value: tokens[0].Value,
	}
	return node, 1, nil
}

func init() {
	Register("string", stringHandler{})
}
