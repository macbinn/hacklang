package parser

import (
	"github.com/macbinn/hacklang/ast"
	"github.com/macbinn/hacklang/token"
)

type callHandler struct {
}

func (callHandler) Parse(tokens []*token.Token) (ast.Node, int, error) {
	if len(tokens) < 3 {
		return nil, 0, ErrSyntaxError
	}

	expr, pos, err := ParseGreedy(tokens,
		"id",
		"function",
		"dot",
	)
	if err != nil {
		return nil, 0, ErrSyntaxError
	}

	matched := false
	var node *ast.CallNode
	for pos < len(tokens) {
		if  tokens[pos].Type != token.LPAREN {
			break
		}
		node = &ast.CallNode{
			Callee:    expr,
			Arguments: nil,
		}

		pos++

		for pos < len(tokens) {
			if tokens[pos].Type == token.RPAREN {
				pos++
				matched = true
				break
			}

			expr, i, err := Parse("expr", tokens[pos:])
			if err != nil {
				return nil, 0, err
			}
			pos += i
			node.Arguments = append(node.Arguments, expr)

			if pos >= len(tokens) {
				return nil, 0, err
			}
			if tokens[pos].Type == token.COMMA {
				pos++
				continue
			}
		}
		expr = node
	}

	if matched {
		return node, pos, nil
	}
	return nil, 0, ErrSyntaxError
}

func init() {
	Register("call", callHandler{})
}
