package parser

import (
	"github.com/macbinn/hacklang/ast"
	"github.com/macbinn/hacklang/token"
)

type functionHandler struct {
}

func (functionHandler) Parse(tokens []*token.Token) (ast.Node, int, error) {
	if len(tokens) < 5 || tokens[0].Type != token.ID ||
		tokens[1].Type != token.EQUAL ||
		tokens[2].Type != token.RARROW ||
		tokens[3].Type != token.LBRACE {
		return nil, 0, ErrSyntaxError
	}
	node := &ast.FunctionNode{
		Arguments: []string{tokens[0].Value},
	}
	pos := 4
	for pos < len(tokens) {
		if tokens[pos].Type == token.RBRACE {
			pos++
			return node, pos, nil
		}

		expr, i, err := Parse("expr", tokens[pos:])
		if err != nil {
			return nil, 0, ErrSyntaxError
		}
		pos += i
		node.Body = append(node.Body, expr)

	}
	return nil, 0, ErrSyntaxError
}

func init() {
	Register("function", functionHandler{})
}
