package parser

import (
	"github.com/macbinn/hacklang/ast"
	"github.com/macbinn/hacklang/token"
)

type idHandler struct {
}

func (idHandler) Parse(tokens []*token.Token) (ast.Node, int, error) {
	if tokens[0].Type == token.ID {
		node := &ast.IdNode{
			Name: tokens[0].Value,
		}
		return node, 1, nil
	}
	return nil, 0, ErrSyntaxError
}

func init() {
	Register("id", idHandler{})
}
