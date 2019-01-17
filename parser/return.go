package parser

import (
	"github.com/macbinn/hacklang/ast"
	"github.com/macbinn/hacklang/token"
)

type returnHandler struct {

}

func (returnHandler) Parse(tokens []*token.Token) (ast.Node, int, error) {
	if tokens[0].Type != token.RETURN {
		return nil, 0, ErrSyntaxError
	}
	expr, pos, err := ParseGreedy(tokens[1:], "expr")
	if err != nil {
		return nil, 0, ErrSyntaxError
	}
	node := &ast.ReturnNode{
		Expr: expr,
	}
	return node, pos + 1, nil
}

func init() {
	Register("return", returnHandler{})
}
