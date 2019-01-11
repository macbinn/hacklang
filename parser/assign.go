package parser

import (
	"github.com/macbinn/hacklang/ast"
	"github.com/macbinn/hacklang/token"
)

type assignHandler struct {
}

func (assignHandler) Parse(tokens []*token.Token) (ast.Node, int, error) {
	if len(tokens) < 3 {
		return nil, 0, ErrSyntaxError
	}
	leftExpr, pos, err := ParseGreedy(tokens, "id", "dot")
	if err != nil {
		return nil, 0, ErrSyntaxError
	}
	if pos >= len(tokens) ||
		tokens[pos].Type != token.EQUAL {
		return nil, 0, ErrSyntaxError
	}
	pos ++
	rightExpr, i, err := Parse("expr", tokens[pos:])
	if err != nil {
		return nil, 0, ErrSyntaxError
	}
	node := &ast.AssignNode{
		Left:  leftExpr,
		Right: rightExpr,
	}
	pos += i
	return node, pos, nil
}

func init() {
	Register("assign", assignHandler{})
}
