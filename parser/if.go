package parser

import (
	"github.com/macbinn/hacklang/ast"
	"github.com/macbinn/hacklang/token"
)

type ifHandler struct {

}

func (ifHandler) Parse(tokens []*token.Token) (ast.Node, int, error) {
	if len(tokens) < 3 {
		return nil, 0, ErrSyntaxError
	}
	if tokens[0].Type != token.IF {
		return nil, 0, ErrSyntaxError
	}
	condExpr, pos, err := ParseGreedy(tokens[1:], "expr")
	if err != nil {
		return nil, 0, ErrSyntaxError
	}
	pos ++
	if tokens[pos].Type != token.LBRACE {
		return nil, 0, ErrSyntaxError
	}
	pos ++
	body, i, err := ParseGreedy(tokens[pos:], "exprList")
	if err != nil {
		return nil, 0, ErrSyntaxError
	}
	pos += i
	if tokens[pos].Type != token.RBRACE {
		return nil, 0, ErrSyntaxError
	}
	node := &ast.IfNode{
		Condition: condExpr,
		Body:      body,
	}
	return node, pos + 1, nil
}

func init() {
	Register("if", ifHandler{})
}
