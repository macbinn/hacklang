package parser

import (
	"github.com/macbinn/hacklang/ast"
	"github.com/macbinn/hacklang/token"
)

type dotHandler struct {
}

func (dotHandler) Parse(tokens []*token.Token) (ast.Node, int, error) {
	if len(tokens) < 3 {
		return nil, 0, ErrSyntaxError
	}
	dotIdx := 0
	for i, t := range tokens{
		if t.Type == token.DOT {
			dotIdx = i
			break
		}
	}
	if dotIdx <= 0 {
		return nil, 0, ErrSyntaxError
	}
	expr, pos, err := ParseGreedy(tokens[:dotIdx],
		"string",
		"list",
		"id",
		"call",
	)
	if err != nil {
		return nil, 0, ErrSyntaxError
	}
	if pos != dotIdx {
		return nil, 0, ErrSyntaxError
	}

	matched := false
	for pos + 1 < len(tokens) {
		if tokens[pos].Type == token.DOT && tokens[pos + 1].Type == token.ID {
			expr = &ast.DotNode{
				Left: expr,
				Right: &ast.IdNode{
					Name: tokens[pos + 1].Value,
				},
			}
			matched = true
			pos += 2
		} else {
			break
		}
	}
	if matched {
		return expr, pos, nil
	}
	return nil, 0, ErrSyntaxError
}

func init() {
	Register("dot", dotHandler{})
}
