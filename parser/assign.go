package parser

import (
	"github.com/macbinn/hacklang/ast"
	"github.com/macbinn/hacklang/token"
)

type assignHandler struct {
}

func (assignHandler) Parse(tokens []*token.Token) (ast.Node, int, error) {
	if len(tokens) < 3 || tokens[0].Type != token.ID || tokens[1].Type != token.EQUAL {
		return nil, 0, ErrSyntaxError
	}
	exprNode, pos, err := Parse("expr", tokens[2:])
	if err != nil {
		return nil, 0, ErrSyntaxError
	}
	node := &ast.AssignNode{
		Left:  &ast.IdNode{Name: tokens[0].Value},
		Right: exprNode,
	}
	return node, pos + 2, nil
}

func init() {
	Register("assign", assignHandler{})
}
