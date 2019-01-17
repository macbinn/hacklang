package parser

import (
	"github.com/macbinn/hacklang/ast"
	"github.com/macbinn/hacklang/token"
)

type notHandler struct {

}

func (notHandler) Parse(tokens []*token.Token) (ast.Node, int, error) {
	if tokens[0].Type != token.NOT {
		return nil, 0, ErrSyntaxError
	}
	expr, pos, err := ParseGreedy(tokens[1:], "expr")
	if err != nil {
		return nil, 0, ErrSyntaxError
	}
	node := &ast.NotNode{
		Expr: expr,
	}
	return node, pos + 1, nil
}

func init() {
	Register("not", notHandler{})
}

