package parser

import (
	"github.com/macbinn/hacklang/ast"
	"github.com/macbinn/hacklang/token"
)

type listHandler struct {
}

func (listHandler) Parse(tokens []*token.Token) (ast.Node, int, error) {

	if len(tokens) < 2 || tokens[0].Type != token.LBRACK {
		return nil, 0, ErrSyntaxError
	}

	node := new(ast.ListNode)
	pos := 1

	for pos < len(tokens) {

		if tokens[pos].Type == token.RBRACK {
			return node, pos + 1, nil
		}

		expr, p, err := Parse("expr", tokens[pos:])
		if err != nil {
			return nil, 0, ErrSyntaxError
		}

		node.Items = append(node.Items, expr)
		pos += p

		if tokens[pos].Type == token.COMMA {
			pos++
			continue
		} else if tokens[pos].Type == token.RBRACK {
			return node, pos + 1, nil
		} else {
			return nil, 0, ErrSyntaxError
		}
	}
	return nil, 0, ErrSyntaxError
}

func init() {
	Register("list", listHandler{})
}
