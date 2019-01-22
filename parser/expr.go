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
		"not",
		"return",
	)
	if err != nil {
		return nil, 0, ErrSyntaxError
	}
	if pos < len(tokens) {
		t := tokens[pos]
		switch t.Type {
		case token.AND:
			pos++
			rightExpr, i, err := ParseGreedy(tokens[pos:], "expr")
			if err != nil {
				return nil, 0, ErrSyntaxError
			}
			pos += i
			node := &ast.AndNode{
				Left:  expr,
				Right: rightExpr,
			}
			return node, pos, nil
		case token.OR:
			pos++
			rightExpr, i, err := ParseGreedy(tokens[pos:], "expr")
			if err != nil {
				return nil, 0, ErrSyntaxError
			}
			pos += i
			node := &ast.OrNode{
				Left:  expr,
				Right: rightExpr,
			}
			return node, pos, nil
		case token.PLUS:
			pos++
			rightExpr, i, err := ParseGreedy(tokens[pos:], "expr")
			if err != nil {
				return nil, 0, ErrSyntaxError
			}
			pos += i
			node := &ast.Add{
				Left:  expr,
				Right: rightExpr,
			}
			return node, pos, nil
		case token.MINS:
			pos++
			rightExpr, i, err := ParseGreedy(tokens[pos:], "expr")
			if err != nil {
				return nil, 0, ErrSyntaxError
			}
			pos += i
			node := &ast.Min{
				Left:  expr,
				Right: rightExpr,
			}
			return node, pos, nil
		case token.MUL:
			pos++
			rightExpr, i, err := ParseGreedy(tokens[pos:], "expr")
			if err != nil {
				return nil, 0, ErrSyntaxError
			}
			pos += i
			node := &ast.Mul{
				Left:  expr,
				Right: rightExpr,
			}
			return node, pos, nil
		case token.DEV:
			pos++
			rightExpr, i, err := ParseGreedy(tokens[pos:], "expr")
			if err != nil {
				return nil, 0, ErrSyntaxError
			}
			pos += i
			node := &ast.Dev{
				Left:  expr,
				Right: rightExpr,
			}
			return node, pos, nil
		case token.EQUALS:
			pos++
			rightExpr, i, err := ParseGreedy(tokens[pos:], "expr")
			if err != nil {
				return nil, 0, ErrSyntaxError
			}
			pos += i
			node := &ast.Equals{
				Left:  expr,
				Right: rightExpr,
			}
			return node, pos, nil
		}
	}
	return expr, pos, nil
}

func init() {
	Register("expr", exprHandler{})
}
